import { Edit, useForm } from "@refinedev/antd";
import { Form, Input } from "antd";

export const OperatorsEdit = () => {
    const { formProps, saveButtonProps } = useForm({});

    return (
        <Edit saveButtonProps={saveButtonProps}>
            <Form {...formProps} layout="vertical">
                <Form.Item
                    label={"name"}
                    name={["name"]}
                    rules={[{required: true}]}
                >
                    <Input />
                </Form.Item>
            </Form>
        </Edit>
    );
};
