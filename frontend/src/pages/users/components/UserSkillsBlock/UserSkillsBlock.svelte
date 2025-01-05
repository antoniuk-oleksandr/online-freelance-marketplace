<script lang="ts">
    import type {Skill} from "@/types/Skill.ts";
    import UserComponentLayout from "@/pages/users/components/UserComponentLayout.svelte";
    import SkillBadge from "@/pages/users/components/UserSkillsBlock/components/SkillBadge/SkillBadge.svelte";

    type UserSkillsBlockProps = {
        skills: Skill[] | null
    }

    let {skills}: UserSkillsBlockProps = $props()
</script>

{#if skills !== null}
    <UserComponentLayout>
        <p class="text-xl font-bold">Skills</p>
        <div class="flex gap-x-3">
            {#each skills as skill }
                <SkillBadge skill={skill}/>
            {/each}
        </div>
    </UserComponentLayout>
{/if}