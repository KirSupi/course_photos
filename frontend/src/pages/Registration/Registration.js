import React, {useEffect, useRef, useState} from 'react';
import axios from "axios";
import {useQuery} from "react-query";
import {
    Text, Card,
    Table, TextInput, Group, Button, Title, Container, Space, PasswordInput, Input
} from "@mantine/core";
import {apiErrorHandler} from "utils/api_errors_handler";
import {useNavigate} from "react-router-dom";
import {reverseSorter, getSorter, sorterById} from "utils/sorters";
import {useForm} from "@mantine/form";

export default function Registration() {
    document.title = 'Регистрация';

    const navigate = useNavigate();
    const phoneRegex = /^(\+?\d{1,3}[-.\s]?)?(\(?\d{3}\)?[-.\s]?)?\d{3}[-.\s]?\d{4}$/;

    const form = useForm({
        mode: 'uncontrolled',
        initialValues: {
            login: '',
            name: '',
            phone: '',
            password: '',
        },
        validate:{
            phone: v=>phoneRegex.test(v) ? null : 'Неверный формат телефона',
        }
    });

    return <Container>
        <Space h={'5rem'}/>
        <Card shadow='sm'
              padding='lg'
              radius='md'
              withBorder>
            <Title>Регистрация</Title>

            <form onSubmit={form.onSubmit((values) => {
                axios
                    .post('/api/register', {
                        login: values.login,
                        password: values.password,
                        name: values.name,
                        phone: values.phone,
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
                <TextInput withAsterisk
                           label="Имя"
                           placeholder="name"
                           key={form.key('name')}
                           {...form.getInputProps('name')}/>
                <TextInput withAsterisk
                       label="Телефон"
                       placeholder="+791234567890"
                       key={form.key('phone')}
                       {...form.getInputProps('phone')}/>
                <PasswordInput withAsterisk
                               label="Пароль"
                               placeholder="password"
                               key={form.key('password')}
                               {...form.getInputProps('password')}/>

                <Group justify="center" mt="md">
                    <Button fullWidth
                            type="submit">
                        Создать аккаунт
                    </Button>
                    <Button fullWidth
                            variant="light"
                            onClick={() => navigate("/login")}>
                        Авторизация
                    </Button>
                </Group>
            </form>
        </Card>
    </Container>
}
