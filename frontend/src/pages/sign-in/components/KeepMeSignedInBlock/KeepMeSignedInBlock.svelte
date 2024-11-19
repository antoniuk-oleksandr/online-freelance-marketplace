<script lang="ts">
    import KeepMeSignedInBlockLayout
        from "@/pages/sign-in/components/KeepMeSignedInBlock/KeepMeSignedInBlockLayout.svelte";
    import KeepMeSignedInBlockIcon
        from "@/pages/sign-in/components/KeepMeSignedInBlock/components/KeepMeSignedInBlockIcon/KeepMeSignedInBlockIcon.svelte";
    import KeepMeSignedInBlockInput
        from "@/pages/sign-in/components/KeepMeSignedInBlock/components/KeepMeSignedInBlockInput/KeepMeSignedInBlockInput.svelte";
    import KeepMeSignedInBlockText
        from "@/pages/sign-in/components/KeepMeSignedInBlock/components/KeepMeSignedInBlockText/KeepMeSignedInBlockText.svelte";

    let keepSignedIn = $state(false);
    const setKeepSignedIn = (value: boolean) => keepSignedIn = value;
</script>

<KeepMeSignedInBlockLayout keepSignedIn={keepSignedIn} setKeepSignedIn={setKeepSignedIn}>
    <KeepMeSignedInBlockIcon keepSignedIn={keepSignedIn}/>
    <KeepMeSignedInBlockInput keepSignedIn={keepSignedIn}/>
    <KeepMeSignedInBlockText/>
</KeepMeSignedInBlockLayout>