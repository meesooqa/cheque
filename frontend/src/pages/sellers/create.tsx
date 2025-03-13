import { Create, useForm } from "@refinedev/antd";
import { Form, Input } from "antd";

export const SellersCreate = () => {
    const { formProps, saveButtonProps } = useForm({});

    return (
        <Create saveButtonProps={saveButtonProps}>
            <Form {...formProps} layout="vertical">
                <Form.Item label={"name"} name={["name"]} rules={[{required: true}]}><Input /></Form.Item>
                <Form.Item label={"inn"} name={["inn"]} rules={[{required: true}]}><Input /></Form.Item>
            </Form>
        </Create>
    );
};
