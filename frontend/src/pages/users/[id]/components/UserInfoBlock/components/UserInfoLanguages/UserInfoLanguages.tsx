import UserInfoLanguagesLayout from "./UserInfoLanguagesLayout";
import {IoLanguage} from "react-icons/io5";
import {UserByIdProps} from "@/types/UserByIdProps";

const UserInfoLanguages = (props: UserByIdProps) => {
    const { user } = props;
    const { languages } = user;

    return (
        <UserInfoLanguagesLayout>
            <IoLanguage/>
            <p className="font-medium">{languages.join(", ")}</p>
        </UserInfoLanguagesLayout>
    )
}

export default UserInfoLanguages;