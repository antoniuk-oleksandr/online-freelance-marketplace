import SkillBadgeLayout from "./SkillBadgeLayout";

type SkillBadgeProps = {
    skill: string
}

const SkillBadge = (props: SkillBadgeProps) => {
    const {skill} = props;

    return (
        <SkillBadgeLayout>
            <span>{skill}</span>
        </SkillBadgeLayout>
    )
}

export default SkillBadge;