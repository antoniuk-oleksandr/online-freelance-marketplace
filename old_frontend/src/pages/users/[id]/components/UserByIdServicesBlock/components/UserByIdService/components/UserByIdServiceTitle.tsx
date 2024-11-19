import {UserService} from "@/types/UserService";

const UserByIdServiceTitle = (props: UserService) => {
    const {title} = props;

    return (
        <p className="text-lg font-semibold mb-2 line-clamp-2 h-14">{title}</p>
    )
}

export default UserByIdServiceTitle;