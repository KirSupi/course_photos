import React, {useEffect, useId, useRef, useState} from 'react';
import axios from "axios";
import {useQuery} from "react-query";
import {apiErrorHandler} from "utils/api_errors_handler";
import {Button, Group, TextInput, Textarea, Flex, Container, FileInput, Image, Text} from "@mantine/core";
import {useForm} from "@mantine/form";
import {notifications} from "@mantine/notifications";
import {uniq} from "lodash";
import {useNavigate} from "react-router-dom";

export default function CreateStudio() {
    document.title = 'Добавление фотостудии';

    const navigate = useNavigate();
    const form = useForm({
        mode: 'uncontrolled',
        initialValues: {
            name: '',
            address: '',
            description: '',
            photos_ids: [],
        },
        validate: {
            name: v => !!v ? null : 'Введите название студии',
            address: v => !!v ? null : 'Введите адрес',
            description: v => !!v ? null : 'Введите описание студии',
            photos_ids: v => !!v && !!v?.length ? null : 'Загрузите фото',
        }
    });

    const [file, setFile] = useState(null);
    useEffect(() => {
        if (!file) return;

        const formData = new FormData();
        formData.append("file", file);

        console.log('axios', file);
        axios
            .post('/api/photo', formData, {
                headers: {
                    'Content-Type': 'multipart/form-data'
                }
            })
            .then(res => {
                if (res.status !== 200) {
                    throw new Error(res.statusText);
                }
                const photo_id = Number(res.data);
                form.setFieldValue('photos_ids', uniq([...(form.getValues()?.photos_ids || []), photo_id]));
                notifications.show({
                    title: 'Фото загружено',
                })
            })
            .catch(apiErrorHandler);

        setFile(null);
    }, [file]);

    return <>
        <Container>
            <form onSubmit={form.onSubmit((values) => {
                axios
                    .post('/api/me/studios', {
                        name: values.name,
                        address: values.address,
                        description: values.description,
                        photos_ids: values.photos_ids,
                    })
                    .then(res => {
                        if (res.status === 200)
                            navigate('/my-studios');

                    })
                    .catch(apiErrorHandler)
            })}>
                <Flex gap='sm' direction='column'>
                    <TextInput withAsterisk
                               label='Название'
                               placeholder='Название студии'
                               key={form.key('name')}
                               {...form.getInputProps('name')}/>
                    <TextInput withAsterisk
                               label='Адрес'
                               placeholder='Адрес студии'
                               key={form.key('address')}
                               {...form.getInputProps('address')}/>
                    <Textarea withAsterisk
                              label='Описание'
                              placeholder='Расскажите о студии'
                              key={form.key('description')}
                              {...form.getInputProps('description')}/>
                    <Flex gap='sm' direction='column'>
                        <Group>
                            <Text>
                                Фотографии студии
                            </Text>
                            <FileInput {...form.getInputProps('photos_ids')}
                                       placeholder='Загрузить (.jpg)'
                                       value={file}
                                       onChange={setFile}
                                       key={form.key('photos_ids')}/>
                        </Group>
                        <Group>
                            {form.getValues()?.photos_ids?.map((id, index) =>
                                <Image src={`/api/photo/${id}`}
                                       h='200px'
                                       fit='contain'
                                       radius='sm'
                                       key={index}/>)}
                        </Group>
                    </Flex>
                    <Group>
                        <Button type='submit'>
                            Создать
                        </Button>
                    </Group>
                </Flex>
            </form>
        </Container>
    </>
}
