import React, {useEffect, useId, useRef, useState} from 'react';
import axios from "axios";
import {useQuery} from "react-query";
import {apiErrorHandler} from "utils/api_errors_handler";
import {Button, Card, Flex, Group, Image, Modal, Space, Text} from "@mantine/core";
import {useNavigate} from "react-router-dom";
import MyStudioBookings from "pages/MyStudios/MyStudioBookings";
import {useDisclosure} from "@mantine/hooks";

export default function MyStudios() {
    document.title = 'Мои фотостудии';

    const navigate = useNavigate();

    const {data: studios, refetch} = useQuery(['/me/studios'], () =>
        axios
            .get('/api/me/studios')
            .then(res => res?.data || [])
            .catch(apiErrorHandler)
    );
    const [selectedStudio, setSelectedStudio] = useState({});
    const [showBookingsOpened, {open: showBookingsOpen, close: showBookingsClose}] = useDisclosure(false);

    return <>
        <Group>
            <Button onClick={() => navigate('/new-studio')}>
                Добавить
            </Button>
        </Group>
        <Space h='lg'/>
        <Flex direction='column' gap='md'>
            {studios?.map((item, index) => (
                <Card key={index}>
                    <Flex direction='column' gap='xs'>
                    <Text fw={500}>{item.name}</Text>
                    <Text size="sm" c="dimmed">{item.address}</Text>
                    <Text size="sm">{item.description}</Text>
                    <Group>
                        {item.photos_ids?.map((photo_id, photo_index) =>
                            <Image src={`/api/photo/${photo_id}`}
                                   h='200px'
                                   fit='contain'
                                   radius='sm'
                                   key={photo_index}/>)}
                    </Group>
                    <Group>
                        <Button variant='light'
                        onClick={()=>{
                            setSelectedStudio(item);
                            showBookingsOpen();
                        }}>
                            Посмотреть брони
                        </Button>
                        <Button variant='light'
                                color='red'>
                            Удалить
                        </Button>
                    </Group>
                    </Flex>
                </Card>
            ))}
        </Flex>
        <Modal opened={showBookingsOpened} onClose={showBookingsClose}
               title={'Бронирования студии ' + selectedStudio?.name}>
            <MyStudioBookings studioId={selectedStudio?.id}/>
        </Modal>
    </>
}
