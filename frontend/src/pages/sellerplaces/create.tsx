import { Create, useForm } from "@refinedev/antd";
import { Form, Input } from "antd";

export const SellerplacesCreate = () => {
    const { formProps, saveButtonProps } = useForm({});
    return (
        <Create saveButtonProps={saveButtonProps}>
            <Form {...formProps} layout="vertical">
                <Form.Item label={"sellerID"} name={["sellerID"]} rules={[{required: true}]}><Input /></Form.Item>
                <Form.Item label={"name"} name={["name"]} rules={[{required: true}]}><Input /></Form.Item>
                <Form.Item label={"address"} name={["address"]} rules={[{required: true}]}><Input /></Form.Item>
                <Form.Item label={"email"} name={["email"]}><Input type="email" /></Form.Item>
            </Form>
        </Create>
    );
};
