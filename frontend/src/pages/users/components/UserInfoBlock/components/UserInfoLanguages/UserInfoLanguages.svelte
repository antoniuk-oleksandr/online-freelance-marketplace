<script lang="ts">
    import UserLanguage from "@/pages/users/components/UserInfoBlock/components/UserLanguage/UserLanguage.svelte";
    import UserInfoLanguagesLayout
        from "@/pages/users/components/UserInfoBlock/components/UserInfoLanguages/UserInfoLanguagesLayout.svelte";
    import Icon from "@iconify/svelte";
    import type {UserByIdProps} from "@/types/UserByIdProps.ts";

    const {user}: UserByIdProps = $props();
</script>

{#if user.languages !== null}
    <UserInfoLanguagesLayout>
        <Icon icon="hugeicons:language-square" />
        {#each user.languages as letter, index}
            <UserLanguage isLast={index === user.languages.length -1} language={letter}/>
        {/each}
    </UserInfoLanguagesLayout>
{/if}