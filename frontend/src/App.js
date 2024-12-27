import React, {useEffect, useState} from 'react';
import {Navigate, Outlet, Route, Routes, useNavigate} from "react-router-dom";
import axios from "axios";
import '@mantine/core/styles.css';
import '@mantine/charts/styles.css';
import '@mantine/dates/styles.css';
import '@mantine/notifications/styles.css';
import {
    AppShell,
    Burger,
    Button,
    Flex,
    Group,
    Loader,
    NavLink,
    ScrollArea,
    Image,
} from "@mantine/core";
import {
    IconUsers,
    IconAdjustments,
    IconCalendarEvent,
    IconPhoto,
    IconLibraryPhoto,
} from "@tabler/icons-react";
import {apiErrorHandler} from "utils/api_errors_handler";
import {useDisclosure} from "@mantine/hooks";
import Catalog from "pages/Catalog/Catalog";
import MyStudios from "pages/MyStudios/MyStudios";
import {useQuery} from "react-query";
import Login from "pages/Login/Login";
import Registration from "pages/Registration/Registration";
import CreateStudio from "pages/MyStudios/CreateStudio";
import MyBookings from "pages/MyBookings/MyBookings";

axios.defaults.withCredentials = true;

function App() {
    const [opened, {toggle}] = useDisclosure();

    const routes = {
        registration: '/registration',
        login: '/login',
        catalog:'/catalog',
        my_studios:'/my-studios',
        new_studio:'/new-studio',
        my_bookings:'/my-bookings',
    }

    const NavItem = ({route, label, Icon}) => {
        const navigate = useNavigate();

        return <NavLink href={route}
                        label={label}
                        leftSection={!!Icon ? <Icon height='1rem'/> : null}
                        active={window.location.pathname === route}
                        onClick={e => {
                            e.preventDefault();
                            navigate(route);
                            toggle();
                        }}/>
    }

    const {data: me, refetch} = useQuery(['/me'], () =>
        axios
            .get('/api/me')
            .then(res => res?.data || null)
            .catch(()=>{})
    );

    return <>
        {!me ? <Flex>
                <Routes>
                    <Route path={routes.login}
                           element={<Login refetchMe={refetch} />}/>
                    <Route path={routes.registration}
                           element={<Registration refetchMe={refetch} />}/>
                    <Route path='*' element={<Navigate to={routes.login}/>}/>
                </Routes>
                <Outlet/>
            </Flex> :
            <AppShell header={{height: '3rem'}}
                      navbar={{
                          width: '14rem',
                          breakpoint: 'md',
                          collapsed: {mobile: !opened},
                      }}
                      padding='md'>
                <AppShell.Header>
                    <Group h='100%' px='md' justify='space-between'>
                        <Group>
                            <Burger opened={opened}
                                    onClick={toggle}
                                    hiddenFrom='sm'
                                    size='sm'/>
                        </Group>
                        <Group>
                            <a href='#' onClick={()=>{
                                axios
                                    .delete('/api/logout')
                                    .then(()=>{
                                        refetch().catch(()=>{});
                                    })
                                    .catch(()=>{
                                        refetch().catch(()=>{});
                                    });
                            }}>{me?.login}</a>
                        </Group>
                    </Group>
                </AppShell.Header>

                <AppShell.Navbar p='md'>
                    <AppShell.Section grow component={ScrollArea}>
                        <NavItem route={routes.catalog}
                                 Icon={IconLibraryPhoto}
                                 label='Каталог фотостудий'/>
                        <NavItem route={routes.my_studios}
                                 Icon={IconPhoto}
                                 label='Мои фотостудии'/>
                        <NavItem route={routes.my_bookings}
                                 Icon={IconCalendarEvent}
                                 label='Мои брони'/>
                    </AppShell.Section>
                </AppShell.Navbar>

                <AppShell.Main>
                    <Routes>
                        <Route path={routes.catalog}
                               element={<Catalog/>}/>
                        <Route path={routes.my_studios}
                               element={<MyStudios/>}/>
                        <Route path={routes.new_studio}
                               element={<CreateStudio/>}/>
                        <Route path={routes.my_bookings}
                               element={<MyBookings/>}/>
                        <Route path='*' element={<Navigate to={routes.catalog}/>}/>
                    </Routes>
                    <Outlet/>
                </AppShell.Main>
            </AppShell>
        }
    </>
}

export default App;
