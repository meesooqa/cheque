import { Show, TextField } from "@refinedev/antd";
import { useShow } from "@refinedev/core";
import { Typography } from "antd";

const { Title } = Typography;

export const SellersShow = () => {
    const { queryResult } = useShow({});
    const { data, isLoading } = queryResult;
    const record = data?.data;
    return (
        <Show isLoading={isLoading}>
            <Title level={5}>{"name"}</Title>
            <TextField value={record?.name} />
            <Title level={5}>{"inn"}</Title>
            <TextField value={record?.inn} />
        </Show>
    );
};
