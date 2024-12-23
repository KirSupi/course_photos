import React, {useEffect, useState} from 'react';
import {notifications} from "@mantine/notifications";
import {urls, constants} from "constants/constants";
import {ActionIcon, CopyButton, rem, Tooltip, Text, Flex} from "@mantine/core";
import {IconCheck, IconCopy} from "@tabler/icons-react";

export const apiErrorHandler = (e) => {
    if (e?.status === 401) {
        if (!constants.developmentMode)
            window.location.replace(urls.auth);
        return;
    }

    // Обработка 500, 501 и т.д.
    if (Math.floor((e?.status || 0) / 100) === 5) {
        notifications.show({
            title: 'Ошибка',
            message: <>
                <Text>При обработке запроса произошла ошибка</Text>
                <br/>
                <Flex align='center'
                      gap='sm'>
                    <Text>trace_id: {e?.response?.data?.trace_id}</Text>
                    <CopyButton value={e?.response?.data?.trace_id + ''} timeout={2000}>
                        {({copied, copy}) => (
                            <Tooltip label={copied ? 'Скопировано' : 'Копировать'} withArrow position="right">
                                <ActionIcon color={copied ? 'teal' : 'gray'} variant="subtle" onClick={copy}>
                                    {copied ? (
                                        <IconCheck style={{width: rem(16)}}/>
                                    ) : (
                                        <IconCopy style={{width: rem(16)}}/>
                                    )}
                                </ActionIcon>
                            </Tooltip>
                        )}
                    </CopyButton>
                </Flex>
            </>
        });
        return;
    }

    notifications.show({
        title: 'Ошибка',
        message: <>
            <Text>При отправке запроса произошла ошибка</Text>
            <br/>
            <Flex align='center'
                  gap='sm'>
                <Text>trace_id: {e?.response?.data?.trace_id}</Text>
                <CopyButton value={e?.response?.data?.trace_id + ''} timeout={2000}>
                    {({copied, copy}) => (
                        <Tooltip label={copied ? 'Скопировано' : 'Копировать'} withArrow position="right">
                            <ActionIcon color={copied ? 'teal' : 'gray'} variant="subtle" onClick={copy}>
                                {copied ? (
                                    <IconCheck style={{width: rem(16)}}/>
                                ) : (
                                    <IconCopy style={{width: rem(16)}}/>
                                )}
                            </ActionIcon>
                        </Tooltip>
                    )}
                </CopyButton>
            </Flex>
        </>
    });

    console.error(e);
}