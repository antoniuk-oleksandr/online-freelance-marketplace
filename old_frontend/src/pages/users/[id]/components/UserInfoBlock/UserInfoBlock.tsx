import UserInfoLayout from "./UserInfoBlockLayout";
import {User} from "@/types/User";
import {Avatar} from "@mantine/core";
import {getFile} from "@/utils/utils";
import UserInfoRightSide
    from "@/pages/users/[id]/components/UserInfoBlock/components/UserInfoRightSide/UserInfoRightSide";
import {useTheme} from "next-themes";
import Link from "next/link";

type UserInfoProps = {
    user: User,
    size?: "small" | "large",
    showLanguages?: boolean,
    useLink?: boolean,
}

const UserInfoBlock = (props: UserInfoProps) => {
    const {user, size, useLink} = props;
    const {firstName, surname, avatar, id} = user;
    const {resolvedTheme} = useTheme();

    const AvatarComponent = useLink ? Link : "div";

    return (
        <UserInfoLayout>
            <AvatarComponent href={`/users/${id}`}>
                <Avatar
                    variant={resolvedTheme}
                    src={avatar && getFile(avatar)}
                    alt={`${firstName} ${surname}`}
                    size={size === "small" ? 64 : 128}
                    color="blue"
                />
            </AvatarComponent>
            <UserInfoRightSide{...props}/>
        </UserInfoLayout>
    );
}

export default UserInfoBlock;
