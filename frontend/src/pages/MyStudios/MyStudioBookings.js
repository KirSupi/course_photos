import React, {useEffect, useId, useRef, useState} from 'react';
import axios from "axios";
import {useQuery} from "react-query";
import {apiErrorHandler} from "utils/api_errors_handler";
import {Button, Card, Flex, Group, Image, Space, Text, Title} from "@mantine/core";
import {useNavigate} from "react-router-dom";
import {notifications} from "@mantine/notifications";
import {formatDate} from "date-fns";
import {formatBookingHours} from "utils/utils";

export default function MyStudioBookings({studio}) {
    const {data: bookings, refetch} = useQuery([`/me/studios/${studio?.id}/bookings`], () =>
        axios
            .get(`/api/me/studios/${studio.id}/bookings`)
            .then(res => res?.data || [])
            .catch(apiErrorHandler)
    );

    return <>
        <Flex direction='column' gap='md'>
            {!!bookings?.length ? bookings?.map((item, index) => (
                <Card key={index}>
                    <Flex direction='column' gap='xs'>
                        <Text fw={500}>{formatDate(new Date(item.date), 'yyyy.MM.dd')} {formatBookingHours(item.hours)}</Text>
                        <Text size="sm">
                            {item.guest_name}, <a href={'tel:' + item.guest_phone}>
                                {item.guest_phone}
                            </a>
                        </Text>
                    </Flex>
                </Card>
            )) : <Text>Бронирований нет</Text>}
        </Flex>
    </>
}
