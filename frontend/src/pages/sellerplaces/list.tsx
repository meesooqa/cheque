import type { BaseRecord } from "@refinedev/core";
import {
    DeleteButton,
    EditButton,
    List,
    ShowButton,
    useTable
} from "@refinedev/antd";
import { Space, Table, Input, Button } from "antd";

export const SellerplacesList = () => {
    const { tableProps } = useTable({});
    return (
        <List>
            <Table {...tableProps} rowKey="id">
                <Table.Column dataIndex="id" title={"id"} />
                <Table.Column dataIndex="name" title={"name"} sorter={true}
                    filterDropdown={({ setSelectedKeys, selectedKeys, confirm, clearFilters }) => (
                    <div style={{ padding: 8 }}>
                        <Input
                          placeholder="Name"
                          value={selectedKeys[0]}
                          onChange={e => setSelectedKeys(e.target.value ? [e.target.value] : [])}
                          onPressEnter={() => confirm()}
                          style={{ marginBottom: 8, display: "block" }}
                        />
                        <Space>
                          <Button onClick={() => {clearFilters && clearFilters();confirm();}}>Reset</Button>
                          <Button type="primary" onClick={() => confirm()}>Search</Button>
                        </Space>
                    </div>
                    )}
                    onFilter={(value, record) => record.name ? record.name.toString().toLowerCase().includes((value as string).toLowerCase()) : false}
                />
                <Table.Column dataIndex="address" title={"address"} sorter={true} />
                <Table.Column dataIndex="email" title={"email"} sorter={true} />
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
