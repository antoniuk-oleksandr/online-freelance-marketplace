import {FaStar} from "react-icons/fa6";
import UserByIdServiceInfoLayout
    from "@/pages/users/[id]/components/UserByIdServicesBlock/components/UserByIdService/components/UserByIdServiceInfo/UserByIdServiceInfoLayout";
import {UserService} from "@/types/UserService";

const UserByIdServiceInfo = (props: UserService) => {
    const {rating, reviewsCount} = props;

    return (
        <UserByIdServiceInfoLayout>
            <FaStar className="text-yellow-400 mr-1  "/>
            <span className="font-semibold vertical-align-middle">{rating}</span>
            <span className="ml-1  vertical-align-middle">({reviewsCount})</span>
        </UserByIdServiceInfoLayout>
    )
}

export default UserByIdServiceInfo;