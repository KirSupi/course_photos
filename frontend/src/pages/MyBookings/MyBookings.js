import React, {useEffect, useId, useRef, useState} from 'react';
import axios from "axios";
import {useQuery} from "react-query";
import {apiErrorHandler} from "utils/api_errors_handler";
import {Button, Card, Flex, Group, Image, Space, Text} from "@mantine/core";
import {useNavigate} from "react-router-dom";
import {notifications} from "@mantine/notifications";
import {formatDate} from "date-fns";
import {formatBookingHours} from "utils/utils";

export default function MyBookings() {
    document.title = 'Мои бронирования';

    const navigate = useNavigate();

    const {data: studios, refetch} = useQuery(['/me/bookings'], () =>
        axios
            .get('/api/me/bookings')
            .then(res => res?.data || [])
            .catch(apiErrorHandler)
    );

    return <>
        <Flex direction='column' gap='md'>
            {studios?.map((item, index) => (
                <Card key={index}>
                    <Flex direction='column' gap='xs'>
                        <Text fw={500}>{item.studio_name}</Text>
                        <Text fw={500}>{formatDate(new Date(item.date), 'yyyy.MM.dd')} {formatBookingHours(item.hours)}</Text>
                        <Text size="sm" c="dimmed">{item.studio_address}</Text>
                        <Text size="sm">
                            {item.owner_name}, <a href={'tel:' + item.owner_phone}>
                                {item.owner_phone}
                            </a>
                        </Text>
                        <Text size="sm">{item.studio_description}</Text>
                        <Group>
                            {item.studio_photos_ids?.map((photo_id, photo_index) =>
                                <Image src={`/api/photo/${photo_id}`}
                                       h='100px'
                                       fit='contain'
                                       radius='sm'
                                       key={photo_index}/>)}
                        </Group>
                        <Group>
                            <Button variant='light'
                                    onClick={() => {
                                        axios
                                            .delete(`/api/me/bookings/${item.id}`)
                                            .then(() => {
                                                notifications.show({
                                                    title: 'Бронь отменена',
                                                });
                                                refetch().catch(apiErrorHandler);
                                            })
                                            .catch(apiErrorHandler)
                                    }}
                                    color='red'>
                                Отменить
                            </Button>
                        </Group>
                    </Flex>
                </Card>
            ))}
        </Flex>
    </>
}
