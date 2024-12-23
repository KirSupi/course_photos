import React, {Fragment, Profiler} from 'react';
import ReactDOM from 'react-dom/client';
import './index.scss';
import App from './App';
import {BrowserRouter} from "react-router-dom";
import {QueryClientProvider, QueryClient} from "react-query";
import {createTheme, MantineProvider} from "@mantine/core";
import {Notifications} from "@mantine/notifications";
import {enableDevelopmentMode} from "constants/constants";
import {profilerCallback} from "utils/profiler";

const queryClient = new QueryClient();

const root = ReactDOM.createRoot(document.getElementById('root'));

if (process.env.NODE_ENV === 'development') {
    enableDevelopmentMode();
}

const theme = createTheme({
    primaryColor: 'green',
});


root.render(
    <React.StrictMode>
            <BrowserRouter>
                <QueryClientProvider client={queryClient}>
                    <MantineProvider theme={theme} forceColorScheme={'dark'}>
                        <Notifications/>
                        <App/>
                    </MantineProvider>
                </QueryClientProvider>
            </BrowserRouter>
    </React.StrictMode>
);
