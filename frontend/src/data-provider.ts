import {
    DataProvider,
    GetListParams,
    GetOneParams
} from "@refinedev/core";
import simpleRestDataProvider from "@refinedev/simple-rest";
import axios from "axios";

const API_URL = import.meta.env.VITE_API_URL;
const baseDataProvider = simpleRestDataProvider(API_URL);

export const dataProvider: DataProvider = {
    ...baseDataProvider,
    getList: async (params: GetListParams) => {
        // pagination
        const { current: page = 1, pageSize: page_size = 10 } = params.pagination ?? {};
        const queryParams = new URLSearchParams({
            page: String(page),
            page_size: String(page_size),
        });
        // sort
        const { sorters } = params;
        if (sorters && sorters.length > 0) {
            queryParams.set(`sort_by`, sorters[0].field);
            queryParams.set(`sort_order`, sorters[0].order);
        } else {
            queryParams.set(`sort_by`, "id");
            queryParams.set(`sort_order`, "desc");
        }
        // filters
        // TODO sum_gt=100&sum_lt=1000&start_date_time=2025-01-01T00:00:00Z&end_date_time=2025-01-31T23:59:59Z
        if (params.filters && params.filters.length > 0) {
            const stringFilters = ["name", "inn"];
            params.filters.forEach((filter) => {
                if ("field" in filter && filter.value) {
                    // const f = filter as { field: string; value: string | number | (string | number)[] };
                    const f = filter as unknown as { field: string; value: string | number };
                    stringFilters.forEach((fieldName: string) => {
                        if (f.field === fieldName) {
                            queryParams.set(fieldName, f.value as string);
                        }
                    });
                }
            });
        }

        const url = `${API_URL}/${params.resource}?${queryParams}`;
        // console.log({url: url});
        const response = await fetch(url);
        const data = await response.json();
        return {
            data: data.items,
            total: data.count,
        };
    },

    getOne: async (params: GetOneParams) => {
        const { resource, id } = params;
        const url = `${API_URL}/${resource}/${id}`;
        const response = await axios.get(url);
        return {
            data: response.data.item,
        };
    },
};
