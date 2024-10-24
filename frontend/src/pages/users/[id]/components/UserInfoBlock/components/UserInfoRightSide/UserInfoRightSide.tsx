import UserInfoRightSideLayout from "./UserInfoRightSideLayout";
import {UserByIdProps} from "@/types/UserByIdProps";
import UserInfoReviews from "@/pages/users/[id]/components/UserInfoBlock/components/UserInfoReviews/UserInfoReviews";
import UserInfoLanguages
    from "@/pages/users/[id]/components/UserInfoBlock/components/UserInfoLanguages/UserInfoLanguages";
import Link from "next/link";

type UserInfoRightSideProps = UserByIdProps & {
    showLanguages?: boolean,
    size?: "small" | "large",
    useLink?: boolean,
}

const UserInfoRightSide = (props: UserInfoRightSideProps) => {
    const {user, showLanguages, size, useLink} = props;
    const {firstName, surname} = user;

    const NameContainer = useLink ? Link : "div";

    return (
        <UserInfoRightSideLayout>
            <NameContainer href={`/users/${user.id}`}>
                <h2
                    className={`font-semibold 
                    ${size === "small" ? "text-lg" : "text-3xl"}
                    ${useLink ? "cursor-pointer hover:underline duration-200 ease-out" : ""}
                    `}
                >{firstName} {surname}</h2>
            </NameContainer>
            <UserInfoReviews {...props}/>
            {showLanguages === true || showLanguages === undefined ?
                <UserInfoLanguages {...props}/>
                : null
            }
        </UserInfoRightSideLayout>
    )
}

export default UserInfoRightSide;