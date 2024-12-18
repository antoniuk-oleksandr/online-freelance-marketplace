<script lang="ts">
    import KeepMeSignedInBlockLayout
        from "@/pages/sign-in/components/KeepMeSignedInBlock/KeepMeSignedInBlockLayout.svelte";
    import KeepMeSignedInBlockInput
        from "@/pages/sign-in/components/KeepMeSignedInBlock/components/KeepMeSignedInBlockInput/KeepMeSignedInBlockInput.svelte";
    import KeepMeSignedInBlockText
        from "@/pages/sign-in/components/KeepMeSignedInBlock/components/KeepMeSignedInBlockText/KeepMeSignedInBlockText.svelte";
    import Checkbox from "@/common-components/Checkbox/Checkbox.svelte";

    let keepSignedIn = $state(false);
    const setKeepSignedIn = (value: boolean) => keepSignedIn = value;
</script>

<KeepMeSignedInBlockLayout keepSignedIn={keepSignedIn} setKeepSignedIn={setKeepSignedIn}>
    <Checkbox selected={keepSignedIn}/>
    <KeepMeSignedInBlockInput keepSignedIn={keepSignedIn}/>
    <KeepMeSignedInBlockText/>
</KeepMeSignedInBlockLayout>