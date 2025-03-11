import { Edit, useForm } from "@refinedev/antd";
import { Form, Input } from "antd";

export const ReceiptsEdit = () => {
    const { formProps, saveButtonProps } = useForm({});

    return (
        <Edit saveButtonProps={saveButtonProps}>
            <Form {...formProps} layout="vertical">
                <Form.Item
                    label={"external_identifier"}
                    name={["external_identifier"]}
                >
                    <Input />
                </Form.Item>
                <Form.Item
                    label={"date_time"}
                    name={["date_time"]}
                    rules={[{required: true}]}
                >
                    <Input />
                </Form.Item>
                <Form.Item
                    label={"sum"}
                    name={["sum"]}
                    rules={[{required: true}]}
                >
                    <Input />
                </Form.Item>
                <Form.Item
                    label={"fiscal_document_number"}
                    name={["fiscal_document_number"]}
                >
                    <Input />
                </Form.Item>
                <Form.Item
                    label={"fiscal_drive_number"}
                    name={["fiscal_drive_number"]}
                >
                    <Input />
                </Form.Item>
                <Form.Item
                    label={"fiscal_sign"}
                    name={["fiscal_sign"]}
                >
                    <Input />
                </Form.Item>
            </Form>
        </Edit>
    );
};
