<script lang="ts">
    import type {UserByIdProps} from "@/types/UserByIdProps.ts";
    import {Link} from "svelte-routing";
    import UserInfoRightSideLayout
        from "@/pages/users/components/UserInfoBlock/components/UserInfoRightSide/UserInfoRightSideLayout.svelte";
    import UserNameLine from "@/pages/users/components/UserInfoBlock/components/UserNameLine/UserNameLine.svelte";
    import UserInfoReviews
        from "@/pages/users/components/UserInfoBlock/components/UserInfoReviews/UserInfoReviews.svelte";
    import UserInfoLanguages
        from "@/pages/users/components/UserInfoBlock/components/UserInfoLanguages/UserInfoLanguages.svelte";

    type UserInfoRightSideProps = UserByIdProps & {
        showLanguages?: boolean,
        size?: "small" | "large",
        useLink?: boolean,
    }

    const {user, showLanguages, useLink, size}: UserInfoRightSideProps = $props();
</script>

<UserInfoRightSideLayout>
    {#if useLink}
        <Link to={`/users/${user.id}`}>
            <UserNameLine
                    useLink={useLink}
                    size={size}
                    user={user}
            />
        </Link>
    {:else}
        <UserNameLine
                useLink={useLink}
                size={size}
                user={user}
        />
    {/if}
    <UserInfoReviews size={size} user={user}/>
    {#if showLanguages === true || showLanguages === undefined}
        <UserInfoLanguages user={user}/>
    {/if}
</UserInfoRightSideLayout>
