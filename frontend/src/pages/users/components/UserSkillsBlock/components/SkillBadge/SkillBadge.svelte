<script lang="ts">
    import type {Skill} from "@/types/Skill";
    import SkillBadgeLayout
        from "@/pages/users/components/UserSkillsBlock/components/SkillBadge/SkillBadgeLayout.svelte";

    type SkillBadgeProps = {
        skill: Skill
    }

    let {skill}: SkillBadgeProps = $props();
</script>

<SkillBadgeLayout id={skill.id}>
    <span>{skill.name}</span>
</SkillBadgeLayout>