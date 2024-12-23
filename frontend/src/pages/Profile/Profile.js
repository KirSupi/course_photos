import React, {useEffect, useRef, useState} from 'react';
import axios from "axios";
import {useQuery} from "react-query";
import {
    Text,
    Table,
} from "@mantine/core";
import {apiErrorHandler} from "utils/api_errors_handler";
import {useNavigate} from "react-router-dom";
import {reverseSorter, getSorter, sorterById} from "utils/sorters";

export default function Profile() {
    document.title = 'Профиль';

    const navigate = useNavigate();

    const {data: presets, isLoading, refetch} = useQuery(['/presets'], () =>
        axios
            .get('/api/presets')
            .then(res => res?.data?.data)
            .catch(apiErrorHandler)
    );

    return <>
        <Table stickyHeader
               highlightOnHover
               withRowBorders>
            <Table.Thead>
                <Table.Tr>
                    <Table.Th>ID</Table.Th>
                    <Table.Th>Название</Table.Th>
                    <Table.Th>Тип</Table.Th>
                    <Table.Th>Описание</Table.Th>
                    <Table.Th visibleFrom={'lg'}>Активен</Table.Th>
                    <Table.Th visibleFrom={'lg'}>Гео</Table.Th>
                </Table.Tr>
            </Table.Thead>
            <Table.Tbody>
                {(presets || [])?.sort(reverseSorter(sorterById))?.map((item, i) =>
                    <Table.Tr key={i}
                              onClick={() => navigate('/presets/' + item.id)}>
                        <Table.Td>
                            {item.id}
                        </Table.Td>
                        <Table.Td>
                            <Text size='sm'>{item.name}</Text>
                        </Table.Td>
                        <Table.Td>{''?.[item?.type] || ''}</Table.Td>
                        <Table.Td>
                            <Text size='sm'>{item.description}</Text>
                        </Table.Td>
                        <Table.Td visibleFrom={'lg'} style={{whiteSpace: 'nowrap'}}>
                            {item?.available_to_use ?
                                <Text size='sm' c='green'>Да</Text> :
                                <Text size='sm' c='dimmed'>Нет</Text>}
                        </Table.Td>
                        <Table.Td visibleFrom={'lg'} style={{whiteSpace: 'nowrap'}}>
                            {(item?.geos || []).join(', ')}
                        </Table.Td>
                    </Table.Tr>)}
            </Table.Tbody>
        </Table>
    </>
}
