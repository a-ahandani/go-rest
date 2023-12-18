import { useMutation, useQueryClient } from "@tanstack/react-query";


export interface User {
    password: string;
    email: string;
}

async function signIn({ email, password }: User) {
    try {
        const response = await fetch('https://localhost:3000/api/auth', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ email, password }),
        });
        const data = await response.json();
        if (!response.ok) {
            throw new Error(data.message || 'Authentication failed');
        }
        return data;
    } catch (error) {
        throw new Error('Authentication failed');
    }
}

const useSIgnIn = () => {
    const queryClient = useQueryClient();
    return useMutation(
        {
            mutationFn: signIn,
            onSuccess: (data: any) => {
                queryClient.setQueryData(['token'], () => data.token);
            },
        }
    );
}
export default useSIgnIn;