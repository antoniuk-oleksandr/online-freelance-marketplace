import UserInfoLanguagesLayout from "./UserInfoLanguagesLayout";
import {IoLanguage} from "react-icons/io5";
import {UserByIdProps} from "@/types/UserByIdProps";

const UserInfoLanguages = (props: UserByIdProps) => {
    const {user} = props;
    const {languages} = user;

    if (!languages) return null;
    return (
        <UserInfoLanguagesLayout>
            <IoLanguage/>
            <p className="font-medium">{languages.map(item => item.name).join(", ")}</p>
        </UserInfoLanguagesLayout>
    )
}

export default UserInfoLanguages;