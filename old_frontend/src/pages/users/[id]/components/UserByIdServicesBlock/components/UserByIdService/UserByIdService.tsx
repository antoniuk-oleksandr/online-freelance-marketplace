import UserByIdServiceLayout from "./UserByIdServiceLayout";
import {UserService} from "@/types/UserService";
import UserByIdServiceImage from "@/pages/users/[id]/components/UserByIdServicesBlock/components/UserByIdService/components/UserByIdServiceImage";
import UserByIdServiceTitle from "@/pages/users/[id]/components/UserByIdServicesBlock/components/UserByIdService/components/UserByIdServiceTitle";
import UserByIdServiceInfo from "@/pages/users/[id]/components/UserByIdServicesBlock/components/UserByIdService/components/UserByIdServiceInfo/UserByIdServiceInfo";
import UserByIdServicePrice from "@/pages/users/[id]/components/UserByIdServicesBlock/components/UserByIdService/components/UserByIdServicePrice";

type UserByIdServiceProps = {
    service: UserService,
}

const UserByIdService = (props: UserByIdServiceProps) => {
    const {service} = props;

    return (
        <UserByIdServiceLayout {...service}>
            <UserByIdServiceImage {...service}/>
            <UserByIdServiceTitle {...service}/>
            <UserByIdServiceInfo {...service}/>
            <UserByIdServicePrice {...service}/>
        </UserByIdServiceLayout>
    )
}

export default UserByIdService;