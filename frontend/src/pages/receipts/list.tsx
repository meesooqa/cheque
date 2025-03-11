import type { BaseRecord } from "@refinedev/core";
import {
    DeleteButton,
    EditButton,
    List,
    ShowButton,
    useTable,
    FilterDropdown
} from "@refinedev/antd";
import { Space, Table, InputNumber, DatePicker } from "antd";
const { RangePicker } = DatePicker;

export const ReceiptsList = () => {
    const { tableProps } = useTable({
        initialSorter: [
            { field: "date_time", order: "desc" },
        ],
    });

    const handleRangePickerChange = (dates: any, dateStrings: any) => {
        console.log("onChange вызван!");
        console.log("Даты:", dates);
        console.log("Строки:", dateStrings);
    };

    return (
        <List>
            <Table {...tableProps} rowKey="externalIdentifier">
                <Table.Column dataIndex="externalIdentifier" title={"externalIdentifier"} />
                <Table.Column
                    dataIndex="dateTime"
                    title={"dateTime"}
                    sorter={true}
                />
                <Table.Column
                    dataIndex="sum"
                    title={"sum"}
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
