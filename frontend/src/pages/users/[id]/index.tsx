import UserPageLayout from "./UserPageLayout";
import {useUserById} from "@/pages/users/[id]/hooks/use-user-by-id";
import AppLoader from "@/common-components/AppLoader";
import UserInfoBlock from "@/pages/users/[id]/components/UserInfoBlock/UserInfoBlock";
import UserAboutBlock from "@/pages/users/[id]/components/UserAboutBlock/UserAboutBlock";
import UserSkillsBlock from "@/pages/users/[id]/components/UserSkillsBlock/UserSkillsBlock";
import UserByIdServicesBlock from "@/pages/users/[id]/components/UserByIdServicesBlock/UserByIdServicesBlock";
import UserByIdReviewBlock from "@/pages/users/[id]/components/UserByIdReviewBlock/UserByIdReviewBlock";
import NotFound from "@/common-components/NotFound/NotFound";

const UserPage = () => {
    const {user, status} = useUserById();

    if (status !== 200 && status !== null) return <NotFound/>;
    if (!user) return <AppLoader/>;
    return (
        <UserPageLayout>
            <UserInfoBlock size={"large"} user={user}/>
            <UserAboutBlock user={user}/>
            <UserSkillsBlock user={user}/>
            <UserByIdServicesBlock user={user}/>
            <UserByIdReviewBlock reviews={user.reviews}/>
        </UserPageLayout>
    )
}

export default UserPage;