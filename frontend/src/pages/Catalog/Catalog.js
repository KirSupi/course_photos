import React, {useState, useEffect} from 'react';
import axios from "axios";
import {useQuery} from "react-query";
import {apiErrorHandler} from "utils/api_errors_handler";
import {Button, Card, Flex, Group, Image, Text, Modal, MultiSelect} from "@mantine/core";
import {useDisclosure} from "@mantine/hooks";
import {useForm} from "@mantine/form";
import {DatePicker, DateTimePicker} from "@mantine/dates";
import {formatDate} from "date-fns";

export default function Catalog() {
    document.title = 'Каталог';

    const {data: studios, refetch} = useQuery(['/studios'], () =>
        axios
            .get('/api/studios')
            .then(res => res?.data || [])
            .catch(apiErrorHandler)
    );
    const [selectedStudio, setSelectedStudio] = useState({});
    const [createBookingOpened, {open: createBookingOpen, close: createBookingClose}] = useDisclosure(false);
    const CreateBooking = () => {
        const form = useForm({
            mode: 'uncontrolled',
            initialValues: {
                date: new Date(),
                hours: [],
            },
        });

        const {data: availableHours, refetch: availableHoursRefetch} = useQuery(['/studios/available-hours'], () =>
            axios
                .get('/api/studios/available-hours', {params: {date: formatDate(form.getValues().date, 'yyyy-MM-dd')}})
                .then(res => res?.data || [])
                .catch(apiErrorHandler)
        );

        useEffect(() => {
            availableHoursRefetch().catch(apiErrorHandler);
        }, [form.values.date]);

        return <>
            <DatePicker label="Логин"
                        placeholder="login"
                        key={form.key('date')}
                        {...form.getInputProps('date')}/>
            <MultiSelect label="Часы аренды"
                         placeholder="Часы аренды"
                         data={availableHours || []}
                         key={form.key('hours')}
                         {...form.getInputProps('hours')}/>
        </>
    }

    return <>
        <Flex direction='column' gap='md'>
            {studios?.map((item, index) => (
                <Card key={index}>
                    <Flex direction='column' gap='xs'>
                        <Text fw={500}>{item.name}</Text>
                        <Text size="sm" c="dimmed">{item.address}</Text>
                        <Text size="sm">{item.owner_name}, <a
                            href={'tel:' + item.owner_phone}>{item.owner_phone}</a></Text>
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
                                    onClick={() => {
                                        setSelectedStudio(item);
                                        createBookingOpen();
                                    }}>
                                Забронировать
                            </Button>
                        </Group>
                    </Flex>
                </Card>
            ))}
            <Modal opened={createBookingOpened} onClose={createBookingClose}
                   title={'Бронирование студии ' + selectedStudio?.name}>
                <CreateBooking/>
            </Modal>
        </Flex>
    </>
}
