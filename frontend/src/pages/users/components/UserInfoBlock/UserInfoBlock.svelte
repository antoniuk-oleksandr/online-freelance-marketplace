<script lang="ts">
    import type {User} from "@/types/User.ts";
    import UserInfoBlockLayout from "@/pages/users/components/UserInfoBlock/UserInfoBlockLayout.svelte";
    import UserInfoRightSide
        from "@/pages/users/components/UserInfoBlock/components/UserInfoRightSide/UserInfoRightSide.svelte";
    import Avatar from "@/common-components/Avatar/Avatar.svelte";

    type UserInfoProps = {
        user: User,
        size?: "small" | "large",
        showLanguages?: boolean,
        useLink?: boolean,
    }

    const {user, size, showLanguages, useLink}: UserInfoProps = $props();
</script>

<UserInfoBlockLayout>
    <Avatar size={size} src={user.avatar} alt="avatar"/>
    <UserInfoRightSide
        useLink={useLink}
        size={size}
        user={user}
        showLanguages={showLanguages}
    />
</UserInfoBlockLayout>