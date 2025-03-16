import React from "react";
import { Tree } from "antd";
import { useList } from "@refinedev/core";

interface Category {
    id: string;
    parentId: string;
    name: string;
}

const CategoryTree: React.FC = () => {
    const buildTree = (items: Category[], parentId = "0"): any => {
        return items
            .filter((item) => item.parentId === parentId)
            .map((item) => ({
                title: item.name,
                key: item.id,
                children: buildTree(items, item.id),
            }));
    };

    const { data, isLoading } = useList<{ items: Category[] }>({
        resource: "categories",
    });
    if (isLoading) {
        return <div>Загрузка...</div>;
    }
    const categories: Category[] = data?.data ?? [];
    const treeData = buildTree(categories);
    return <Tree treeData={treeData} defaultExpandAll />;
};

export default CategoryTree;
