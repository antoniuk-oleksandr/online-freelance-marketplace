import SkillBadgeLayout from "./SkillBadgeLayout";
import {Skill} from "@/types/Skill";

type SkillBadgeProps = {
    skill: Skill
}

const SkillBadge = (props: SkillBadgeProps) => {
    const {skill} = props;
    const {name, id} = skill;

    return (
        <SkillBadgeLayout id={id}>
            <span>{name}</span>
        </SkillBadgeLayout>
    )
}

export default SkillBadge;