import React, { useState } from "react";
import { Edit, useForm, useSelect } from "@refinedev/antd";
import { UploadOutlined, PlusOutlined } from "@ant-design/icons";
import { message, Modal, List, Card, Form, Input, Select, Upload, Button } from "antd";
import type { UploadFile, UploadProps } from "antd/es/upload/interface";
import { useApiUrl } from "@refinedev/core";

export const ProductsEdit: React.FC = () => {
    const apiUrl = useApiUrl();
    const [fileList, setFileList] = useState<UploadFile[]>([]);
    const [previewOpen, setPreviewOpen] = useState(false);
    const [previewImage, setPreviewImage] = useState('');
    const [previewTitle, setPreviewTitle] = useState('');

    const { formProps, saveButtonProps, queryResult } = useForm({
        metaData: {
            fields: ["id", "name", "brandId", "categories",
                "images.id", "images.name", "images.url", "images.isMain", "images.order"],
        },
    });

    const { selectProps: brandSelectProps } = useSelect({
        resource: "brands",
        optionLabel: "name",
        optionValue: "id",
    });

    const { selectProps: categorySelectProps } = useSelect({
        resource: "categories",
        optionLabel: "name",
        optionValue: "id",
    });

    // Загрузка текущих изображений при загрузке формы
    React.useEffect(() => {
        if (queryResult?.data?.data) {
            const product = queryResult.data.data;
            if (product.images) {
                const currentImages = product.images.map((image: any) => ({
                    uid: image.id.toString(),
                    name: image.name,
                    status: 'done',
                    url: image.url,
                    isMain: image.isMain,
                    order: image.order,
                    // Сохраняем ID изображения для последующего использования
                    imageId: image.id
                }));
                setFileList(currentImages);
            }
        }
    }, [queryResult?.data?.data]);

    const handlePreview = async (file: UploadFile) => {
        setPreviewImage(file.url || (file.preview as string));
        setPreviewOpen(true);
        setPreviewTitle(file.name || file.url!.substring(file.url!.lastIndexOf('/') + 1));
    };

    const handleChange: UploadProps['onChange'] = ({ fileList: newFileList }) => {
        setFileList(newFileList);
    };

    const handleSetMainImage = (uid: string) => {
        setFileList(prev =>
            prev.map(file => ({
                ...file,
                isMain: file.uid === uid
            }))
        );
    };

    const uploadProps: UploadProps = {
        action: `${apiUrl}/api/v1/images`, // Использование вашего API для изображений
        listType: "picture-card",
        fileList,
        onPreview: handlePreview,
        onChange: handleChange,
        beforeUpload: (file) => {
            const isImage = file.type.startsWith('image/');
            if (!isImage) {
                message.error('Вы можете загружать только изображения!');
            }
            return isImage || Upload.LIST_IGNORE;
        },
        // Настройка обработки ответа сервера после загрузки
        onSuccess: (response, file, fileList) => {
            // Предполагаем, что сервер возвращает ID созданного изображения
            // и другие необходимые данные
            if (response && response.id) {
                const targetFile = fileList.find(item => item.uid === file.uid);
                if (targetFile) {
                    targetFile.imageId = response.id;
                    targetFile.url = response.url || targetFile.url;
                }
            }
        },
    };

    // Кастомизация формы для включения изображений
    const customFormProps = {
        ...formProps,
        onFinish: async (values: any) => {
            // categories
            if (values.categories && Array.isArray(values.categories)) {
                values.categories = values.categories.map((id: string) => ({ id }));
            }
            // Преобразуем список загруженных файлов в формат для API
            const images = fileList.map((file, index) => ({
                id: file.imageId || (file.uid.startsWith('rc-upload') ? undefined : Number(file.uid)),
                name: file.name,
                url: file.url || (file.response && file.response.url),
                isMain: !!file.isMain,
                order: index,
            })).filter(img => img.id || img.url); // Фильтруем только действительные изображения

            // Объединяем основные данные и данные изображений
            const dataToSubmit = {
                ...values,
                images,
            };

            // Передаем данные в оригинальный обработчик
            return formProps.onFinish && formProps.onFinish(dataToSubmit);
        },
    };

    return (
        <Edit saveButtonProps={saveButtonProps}>
            <Form {...customFormProps} layout="vertical">
                <Form.Item
                    label="Название продукта"
                    name="name"
                    rules={[{ required: true, message: 'Пожалуйста, введите название продукта' }]}
                >
                    <Input />
                </Form.Item>

                <Form.Item
                    label="Бренд"
                    name="brandId"
                >
                    <Select {...brandSelectProps} allowClear />
                </Form.Item>

                <Form.Item
                    label="Категории"
                    name={["categories"]}
                >
                    <Select
                        {...categorySelectProps}
                        mode="multiple"
                        allowClear
                    />
                </Form.Item>

                <Form.Item label="Изображения">
                    <Upload {...uploadProps}>
                        <div>
                            <PlusOutlined />
                            <div style={{ marginTop: 8 }}>Загрузить</div>
                        </div>
                    </Upload>

                    <Modal
                        open={previewOpen}
                        title={previewTitle}
                        footer={null}
                        onCancel={() => setPreviewOpen(false)}
                    >
                        <img alt="preview" style={{ width: '100%' }} src={previewImage} />
                    </Modal>

                    {fileList.length > 0 && (
                        <List
                            header={<div>Управление изображениями</div>}
                            bordered
                            dataSource={fileList}
                            renderItem={(item) => (
                                <List.Item
                                    actions={[
                                        <Button
                                            type={item.isMain ? "primary" : "default"}
                                            onClick={() => handleSetMainImage(item.uid)}
                                        >
                                            {item.isMain ? "Основное изображение" : "Сделать основным"}
                                        </Button>
                                    ]}
                                >
                                    <List.Item.Meta
                                        avatar={<img src={item.url || item.thumbUrl} width={50} alt={item.name} />}
                                        title={item.name}
                                    />
                                </List.Item>
                            )}
                        />
                    )}
                </Form.Item>
            </Form>
        </Edit>
    );
};