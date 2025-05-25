export default function PrivacyPage() {
  const padding_bottom = "pb-8";

  return (
    <div className="container mx-auto p-4 text-left">
      <h1 className={padding_bottom}>Privacy Policy</h1>

      <p className={padding_bottom}>
        A username and e-mail address are collected at the time you register for this service. Along with that information, you
        may be asked to provide other things, including but not limited too your name, age, and various other contact addresses.
        Only a fingerprint of your e-mail address and password are stored. Any other sensitive or personal information is stored but
        not displayed on The Site to other users without your consent.</p>

      <p className={padding_bottom}>
        The Site will set and access cookies on your computer. The information stored therein is used for managing your access to
        certain areas on the site and establishing your identity.</p>

      <p className={padding_bottom}>
        The Site may log any information from an incoming request, including, but not limited to, the client IP address, referring
        URL, and/or user agent string.</p>

      <p className={padding_bottom}>
        The Site does not rent, sell, or share personal information about you with other people or non-affiliated companies except
        to provide services you have requested, when we have your permission, or in response to a court order or subpoena.</p>

    </div>
  );
}