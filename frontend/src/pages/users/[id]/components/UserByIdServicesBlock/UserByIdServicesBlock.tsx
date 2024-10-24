import UserComponentLayout from "@/pages/users/[id]/components/UserComponentLayout";
import { UserByIdProps } from "@/types/UserByIdProps";
import UserByIdService from "@/pages/users/[id]/components/UserByIdServicesBlock/components/UserByIdService/UserByIdService";

const UserByIdServicesBlock = ({ user }: UserByIdProps) => {
    const { services } = user;

    if (!services || services.length === 0) return null;

    return (
        <UserComponentLayout>
            <h2 className="text-xl font-bold">Services</h2>
            <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-4">
                {services.map((service, index) => (
                    <UserByIdService key={index} service={service} />
                ))}
            </div>
        </UserComponentLayout>
    );
}

export default UserByIdServicesBlock;
