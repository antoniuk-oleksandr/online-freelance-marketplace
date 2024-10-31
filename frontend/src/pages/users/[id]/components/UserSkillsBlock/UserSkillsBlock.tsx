import {UserByIdProps} from "@/types/UserByIdProps";
import SkillBadge from "@/pages/users/[id]/components/UserSkillsBlock/components/SkillBadge/SkillBadge";
import UserComponentLayout from "@/pages/users/[id]/components/UserComponentLayout";

const UserSkillsBlock = (props: UserByIdProps) => {
    const {user} = props;
    const {skills} = user;

    if (!skills) return null;
    return (
        <UserComponentLayout>
            <p className={"text-xl font-bold"}>Skills</p>
            <div className={"flex gap-x-2"}>
                {skills.map((skill, index) => (
                    <SkillBadge skill={skill} key={index}/>
                ))}
            </div>
        </UserComponentLayout>
    )
}

export default UserSkillsBlock;