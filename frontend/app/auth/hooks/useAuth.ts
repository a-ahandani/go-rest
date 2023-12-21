import { useMutation, useQueryClient } from "@tanstack/react-query";
import { useRouter } from 'next/navigation'


export interface User {
    password: string;
    email: string;
}

async function signIn({ email, password }: User) {
    try {
        const response = await fetch('http://localhost:3000/api/auth', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ email, password }),
        });
        const data = await response.json();
        if (!response.ok) {
            throw new Error(data.message);
        }
        return data;
    } catch (error) {
        const { message } = error as Error;
        throw new Error(message);
    }
}

const useSignIn = () => {
    const queryClient = useQueryClient();
    const router = useRouter()

    return useMutation(
        {
            mutationFn: signIn,
            onSuccess: (data: any) => {
                queryClient.setQueryData(['token'], () => data.token);
                router.push('/users')
            },
        }
    );
}
export default useSignIn;