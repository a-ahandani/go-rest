import { useQuery, useQueryClient } from "@tanstack/react-query";


async function fetchUsers(token: string) {
    try {
        const response = await fetch("http://localhost:3000/api/users", {
            method: "GET",
            headers: {
                Authorization: `Bearer ${token}`,
            },
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


const useUsers = () => {
    const queryClient = useQueryClient();
    const token = queryClient.getQueryData<string>(["token"]);
    console.log({ token });

    return useQuery({
        queryKey: ["users"],
        queryFn: () => fetchUsers(token || ""),
        enabled: !!token,
    });
};
export default useUsers;