'use client';
import './globals.css'
import {
    HydrationBoundary,
    QueryClient,
    QueryClientProvider,
    dehydrate
} from '@tanstack/react-query'
import React from 'react'

export default function Data({ children }: {
    children: React.ReactNode
}) {
    const [queryClient] = React.useState(() => new QueryClient())
    const dehydratedState = dehydrate(queryClient);

    return (
        <QueryClientProvider client={queryClient}>
            <HydrationBoundary state={dehydratedState}>
                {children}
            </HydrationBoundary>
        </QueryClientProvider>
    )
}