'use client';

import useUsers from "./hooks/useUsers";

const Users = () => {

    const { data } = useUsers()
    console.log(data)
    return <div className="flex items-center justify-center min-h-screen">

        <table className="table">
            {/* head */}
            <thead>
                <tr>
                    <th>
                        <label>
                            <input type="checkbox" className="checkbox" />
                        </label>
                    </th>
                    <th>Name</th>
                    <th>Job</th>
                    <th>Favorite Color</th>
                    <th></th>
                </tr>
            </thead>
            <tbody>
                {/* row 1 */}
                {data?.data?.map((user: any) => {
                    return <tr key={user.Id}>
                        <th>
                            <label>
                                <input type="checkbox" className="checkbox" />
                            </label>
                        </th>
                        <td>
                            <div className="flex items-center gap-3">
                                <div className="font-bold">{user.Name}</div>
                            </div>
                        </td>
                        <td>
                            {user.Email}
                        </td>
                        <th>
                            <button className="btn btn-ghost btn-xs">details</button>
                        </th>
                    </tr>
                })
                }

            </tbody>
            {/* foot */}
            <tfoot>
                <tr>
                    <th></th>
                    <th>Name</th>
                    <th>Job</th>
                    <th>Favorite Color</th>
                    <th></th>
                </tr>
            </tfoot>

        </table>
    </div >
}
export default Users;