import UserByIdReviewUserLayout from "./UserByIdReviewUserLayout";
import {Avatar} from "@mantine/core";
import {getFile} from "@/utils/utils";
import {User} from "@/types/User";
import Link from "next/link";

type UserByIdReviewUserProps = User;

const UserByIdReviewUser = (props: UserByIdReviewUserProps) => {
    const {avatar, firstName, surname, id} = props;
    const link = `/users/${id}`;

    return (
        <UserByIdReviewUserLayout>
            <Link href={link}>
                <Avatar src={avatar && getFile(avatar)} size={48}/>
            </Link>
            <Link href={link}>
                <p className={"hover:underline ease-out duration-200"}>{firstName} {surname}</p>
            </Link>
        </UserByIdReviewUserLayout>
    )
}

export default UserByIdReviewUser;