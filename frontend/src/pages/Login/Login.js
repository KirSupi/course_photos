import React, {useEffect, useRef, useState} from 'react';
import axios from "axios";
import {useQuery} from "react-query";
import {
    Text, Card,
    Table, TextInput, Group, Button, Title, Container, Space, PasswordInput
} from "@mantine/core";
import {apiErrorHandler} from "utils/api_errors_handler";
import {useNavigate} from "react-router-dom";
import {reverseSorter, getSorter, sorterById} from "utils/sorters";
import {useForm} from "@mantine/form";

export default function Login({refetchMe}) {
    document.title = 'Авторизация';

    const navigate = useNavigate();
    const form = useForm({
        mode: 'uncontrolled',
        initialValues: {
            login: '',
            password: '',
        },
    });

    return <Container>
        <Space h={'5rem'}/>
        <Card shadow='sm'
              padding='lg'
              radius='md'
              withBorder>
            <Title>Авторизация</Title>

            <form onSubmit={form.onSubmit((values) => {
                axios
                    .post('/api/login', {
                        login: values.login,
                        password: values.password,
                    })
                    .then(res => {
                        if (res.status === 200)
                            refetchMe().catch(apiErrorHandler);
                    })
                    .catch(apiErrorHandler)
            })}>
                <TextInput withAsterisk
                           label="Логин"
                           placeholder="login"
                           key={form.key('login')}
                           {...form.getInputProps('login')}/>
                <PasswordInput withAsterisk
                               label="Пароль"
                               placeholder="password"
                               key={form.key('password')}
                               {...form.getInputProps('password')}/>

                <Group justify="center" mt="md">
                    <Button fullWidth
                            type="submit">
                        Войти
                    </Button>
                    <Button fullWidth
                            variant="light"
                    onClick={()=>navigate("/registration")}>
                        Регистрация
                    </Button>
                </Group>
            </form>
        </Card>
    </Container>

}
