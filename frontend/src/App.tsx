import { GitHubBanner, Refine } from "@refinedev/core";
import { DevtoolsPanel, DevtoolsProvider } from "@refinedev/devtools";
import { RefineKbar, RefineKbarProvider } from "@refinedev/kbar";

import {
    ErrorComponent,
    ThemedLayoutV2,
    ThemedSiderV2,
    useNotificationProvider,
} from "@refinedev/antd";
import "@refinedev/antd/dist/reset.css";

import routerBindings, {
    DocumentTitleHandler,
    NavigateToResource,
    UnsavedChangesNotifier,
} from "@refinedev/react-router";
import { dataProvider } from "./data-provider";
import { App as AntdApp } from "antd";
import { BrowserRouter, Outlet, Route, Routes } from "react-router";
import { Header } from "./components/header";
import { ColorModeContextProvider } from "./contexts/color-mode";
import {ReceiptsCreate, ReceiptsEdit, ReceiptsList, ReceiptsShow} from "./pages/receipts";
import {OperatorsCreate, OperatorsEdit, OperatorsList, OperatorsShow} from "./pages/operators";

function App() {
    return (
        <BrowserRouter>
            <GitHubBanner />
            <RefineKbarProvider>
                <ColorModeContextProvider>
                    <AntdApp>
                        <DevtoolsProvider>
                            <Refine
                                dataProvider={dataProvider}
                                notificationProvider={useNotificationProvider}
                                routerProvider={routerBindings}
                                resources={[
                                    {
                                        name: "receipts",
                                        list: "/receipts",
                                        create: "/receipts/create",
                                        edit: "/receipts/edit/:id",
                                        show: "/receipts/show/:id",
                                        meta: {
                                            canDelete: true,
                                        },
                                    },
                                    {
                                        name: "operators",
                                        list: "/operators",
                                        create: "/operators/create",
                                        edit: "/operators/edit/:id",
                                        show: "/operators/show/:id",
                                        meta: {
                                            canDelete: true,
                                        },
                                    },
                                ]}
                                options={{
                                    syncWithLocation: true,
                                    warnWhenUnsavedChanges: true,
                                    useNewQueryKeys: true,
                                    projectId: "A6687B38-9D96-4487-BB87-CB6291D2209D",
                                }}
                            >
                                <Routes>
                                    <Route
                                        element={
                                            <ThemedLayoutV2
                                                Header={() => <Header sticky />}
                                                Sider={(props) => <ThemedSiderV2 {...props} fixed />}
                                            >
                                                <Outlet />
                                            </ThemedLayoutV2>
                                        }
                                    >
                                        <Route
                                            index
                                            element={<NavigateToResource resource="receipts" />}
                                        />
                                        <Route path="/receipts">
                                            <Route index element={<ReceiptsList />} />
                                            <Route path="create" element={<ReceiptsCreate />} />
                                            <Route path="edit/:id" element={<ReceiptsEdit />} />
                                            <Route path="show/:id" element={<ReceiptsShow />} />
                                        </Route>
                                        <Route path="/operators">
                                            <Route index element={<OperatorsList />} />
                                            <Route path="create" element={<OperatorsCreate />} />
                                            <Route path="edit/:id" element={<OperatorsEdit />} />
                                            <Route path="show/:id" element={<OperatorsShow />} />
                                        </Route>
                                        <Route path="*" element={<ErrorComponent />} />
                                    </Route>
                                </Routes>
                                <RefineKbar />
                                <UnsavedChangesNotifier />
                                <DocumentTitleHandler />
                            </Refine>
                            <DevtoolsPanel />
                        </DevtoolsProvider>
                    </AntdApp>
                </ColorModeContextProvider>
            </RefineKbarProvider>
        </BrowserRouter>
    );
}

export default App;
