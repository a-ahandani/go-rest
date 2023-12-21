'use client';
import './globals.css'
import { PersistQueryClientProvider } from '@tanstack/react-query-persist-client'
import { createSyncStoragePersister } from '@tanstack/query-sync-storage-persister'

import {
    HydrationBoundary,
    QueryClient,
    dehydrate
} from '@tanstack/react-query'
import React from 'react'

export default function Data({ children }: {
    children: React.ReactNode
}) {

    const queryClient = new QueryClient({
        defaultOptions: {
            queries: {
                gcTime: 1000 * 60 * 60 * 24, // 24 hours
            },
        },
    })
    const dehydratedState = dehydrate(queryClient);

    const persister = createSyncStoragePersister({
        storage: window.localStorage,
    })
    return (
        <PersistQueryClientProvider
            client={queryClient}
            persistOptions={{ persister }}
        >
            <HydrationBoundary state={dehydratedState}>
                {children}
            </HydrationBoundary>
        </PersistQueryClientProvider>
    )
}