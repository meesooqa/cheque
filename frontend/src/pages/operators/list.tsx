import type { BaseRecord } from "@refinedev/core";
import {
    DeleteButton,
    EditButton,
    List,
    ShowButton,
    useTable
} from "@refinedev/antd";
import { Space, Table } from "antd";

export const OperatorsList = () => {
    const { tableProps } = useTable({
        initialSorter: [
            { field: "id", order: "desc" },
        ],
    });

    return (
        <List>
            <Table {...tableProps} rowKey="id">
                <Table.Column dataIndex="id" title={"id"} />
                <Table.Column
                    dataIndex="name"
                    title={"name"}
                    sorter={true}
                />
                <Table.Column
                    title={"Actions"}
                    dataIndex="actions"
                    render={(_, record: BaseRecord) => (
                        <Space>
                            <EditButton hideText size="small" recordItemId={record.id} />
                            <ShowButton hideText size="small" recordItemId={record.id} />
                            <DeleteButton hideText size="small" recordItemId={record.id} />
                        </Space>
                    )}
                />
            </Table>
        </List>
    );
};
