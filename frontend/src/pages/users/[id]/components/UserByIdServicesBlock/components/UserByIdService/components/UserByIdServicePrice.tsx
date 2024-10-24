import {UserService} from "@/types/UserService";

const UserByIdServicePrice = (props: UserService) => {
    const {minPrice} = props;

    return (
        <p className="text-base font-medium">From ${minPrice}</p>
    )
}

export default UserByIdServicePrice;