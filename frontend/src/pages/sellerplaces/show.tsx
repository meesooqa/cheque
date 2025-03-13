import { Show, TextField } from "@refinedev/antd";
import { useShow } from "@refinedev/core";
import { Typography } from "antd";

const { Title } = Typography;

export const SellerplacesShow = () => {
    const { queryResult } = useShow({});
    const { data, isLoading } = queryResult;
    const record = data?.data;
    return (
        <Show isLoading={isLoading}>
            <Title level={5}>{"id"}</Title>
            <TextField value={record?.id} />
            <Title level={5}>{"sellerID"}</Title>
            <TextField value={record?.sellerID} />
            <Title level={5}>{"name"}</Title>
            <TextField value={record?.name} />
            <Title level={5}>{"address"}</Title>
            <TextField value={record?.address} />
            <Title level={5}>{"email"}</Title>
            <TextField value={record?.email} />
        </Show>
    );
};
