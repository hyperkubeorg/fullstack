export default function TermsPage() {
  const padding_bottom = "pb-8";

  return (
    <div className={"container mx-auto p-4 text-left " + padding_bottom}>
      <h1 className={padding_bottom}>Terms of Service</h1>
      <h2 className={padding_bottom}>Definitions</h2>
      "The Site" is defined as this website and its affiliates. "Site Staff" are defined as any individual or group of individuals 
      acting as an Officer of the site to moderate the forums, private messages, and other areas of public discourse.

      <h2 className={padding_bottom}>Description of Service</h2>
      <p className={padding_bottom}>
        The Site provides users with access to a collection of writings, images, videos, and other content through its network 
        of properties which may be accessed, uploaded, and modified through any various medium or device. You also understand 
        and agree that the service may include advertisements and that these advertisements are required to maintain the service. 
        You understand and agree that the service is provided "as is" and The Site assumes no responsibilities for the timelines, 
        deletion, miss-delivery, or failure to store any resource.</p>

      <h2 className={padding_bottom}>Termination</h2>

      <p className={padding_bottom}>
        You understand and agree that Site Staff may terminate access your account if you repeatedly violate these terms of service 
        or if there are unexpected technical or security problems.</p>

      <h2 className={padding_bottom}>Member Conduct</h2>
      <p className={padding_bottom}>
        You understand that all writings, images, videos, audio clips, and other resources transmitted are the sole responsibility 
        of the person from whom such content originated. The Site does not control the content provided by this service and does 
        not guarantee its accuracy, integrity, or quality. You may be exposed to content which you deem to be indecent, objectionable, 
        or otherwise offensive. While Site Staff make an effort to remove material that may be deemed offensive or illegal, under no 
        circumstance will The Site or Site Staff be liable for such content.</p>

      <p className={padding_bottom}>
        Furthermore, you agree not to use the service to:</p>
      <ul>
        <li>upload, post, or otherwise transmit any content that is unlawful in the United States or under international law,</li>
        <li>upload, post, or otherwise transmit works protected by copyright or trademark except in cases of fair use, parody, 
          or other exceptions provided by relevant laws,</li>
        <li>upload, post, or otherwise transmit "spam" material,</li>
        <li>attempt to access another user's account or impersonate a user without their consent,</li>
        <li>attempt to circumvent any protection system or "hack" The Site or a user.</li>
      </ul>

      <p className={padding_bottom}>
        You acknowledge that The Site and Site Staff do not prescreen content and cannot be held liable for materials of a 
        prohibited nature.</p>

      <h2 className={padding_bottom}>Content Submitted to the Service</h2>

      <p className={padding_bottom}>
        The Site does not claim ownership over any material uploaded, posted, or otherwise transmitted to the site by a user. 
        However, by uploading, submitting, or otherwise transmitting content to The Site for inclusion in the service, you are 
        granting us a limited right to display, perform, and/or reproduce this content, revokable by deleting the content from 
        the service or requesting that it be removed in cases where you do not have access to a delete function.</p>

      <p className={padding_bottom}>
        You understand that your content may be displayed besides advertisements to benefit The Site. Due to the nature of The 
        Site' advertising partner, you also understand that said ads may analyze your content and display advertisements 
        relevant to your content, and that a competitor to your content may be shown. The Site is unable to control this and 
        you agree that we cannot be held liable for it.</p>

      <h2 className={padding_bottom}>No Warranty</h2>
      <p className={padding_bottom}>
        Because this service is provided free of charge, there is no warranty for the service to the extent permitted by applicable 
        laws. This service is provided "as is" without warranty of any kind, either express or implied, including, but not limited 
        to, the implied warranties of merchantability, fitness for a particular purpose and non-infringement.</p>

      <p className={padding_bottom}>
        Any material downloaded or otherwise obtained through use of this service is accessed at your own risk, and you will be solely 
        responsible for any damage to your computer system or loss of such data that results from the download of any such material.</p>

      <p className={padding_bottom}>
        No advise or information obtained by you from The Site shall create any warranty not expressly stated in this terms of 
        service document.</p>

      <h2 className={padding_bottom}>Limitation of Liability</h2>
      <p className={padding_bottom}>
        You understand and agree that Site Staff and The Site, its officers, employees, and partners shall not be liable to you for 
        any direct, indirect, incidental, special, consequential, or exemplar damages, including, but not limited to, damages for 
        loss of profits, goodwill, use, data or other intangible losses resulting from the inability to use the service, unauthorized 
        access to or alteration of your transmission or data, statements or conduct of any third-party on the service, or any other 
        matter relating to the service.</p>

      <h2 className={padding_bottom}>Changes to this Agreement</h2>
      <p className={padding_bottom}>
        The Site reserves the right to modify the terms of this agreement at any time. You are responsible for reviewing these terms 
        periodically to ensure you are aware of any changes.</p>

      <h2 className={padding_bottom}>Choice of Law and Forum</h2>
      <p className={padding_bottom}>
        This terms of service document and the relationship between you and The Site will be governed by the laws of the State of Texas
        without regard to its conflict of law provisions. You and The Site agree to submit to the personal and exclusive jurisdiction of
        the courts located within the county of New Haven, Connecticut.</p>
    </div>
  );
}