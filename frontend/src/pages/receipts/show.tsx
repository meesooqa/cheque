import { Show, TextField } from "@refinedev/antd";
import { useShow } from "@refinedev/core";
import { Typography } from "antd";

const { Title } = Typography;

// GET http://localhost:8080/api/v1/receipts/1979
// DELETE http://localhost:8080/api/v1/receipts/1979
export const ReceiptsShow = () => {
    const { queryResult } = useShow({});
    const { data, isLoading } = queryResult;

    const record = data?.data;

    return (
        <Show isLoading={isLoading}>
            <Title level={5}>{"externalIdentifier"}</Title>
            <TextField value={record?.externalIdentifier} />
            <Title level={5}>{"dateTime"}</Title>
            <TextField value={record?.dateTime} />
            <Title level={5}>{"sum"}</Title>
            <TextField value={record?.sum} />
        </Show>
    );
};
