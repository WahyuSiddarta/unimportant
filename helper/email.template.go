package helper

var (
	// TemplateEmailVerify : Email Verify
	TemplateEmailVerify = `<!doctype html>
	<html xmlns=http://www.w3.org/1999/xhtml xmlns:v=urn:schemas-microsoft-com:vml xmlns:o=urn:schemas-microsoft-com:office:office>
	<head>
	<!--[if gte mso 15]>
	<xml>
	<o:OfficeDocumentSettings>
	<o:AllowPNG/>
	<o:PixelsPerInch>96</o:PixelsPerInch>
	</o:OfficeDocumentSettings>
	</xml>
	<![endif]-->
	<meta charset=UTF-8>
	<meta http-equiv=X-UA-Compatible content="IE=edge">
	<meta name=viewport content="width=device-width, initial-scale=1">
	<title>Ruang Nyaman</title>
	<style type=text/css>p{margin:10px 0;padding:0}table{border-collapse:collapse}h1,h2,h3,h4,h5,h6{display:block;margin:0;padding:0}img,a img{border:0;height:auto;outline:0;text-decoration:none}body,#bodyTable,#bodyCell{height:100%;margin:0;padding:0;width:100%}.mcnPreviewText{display:none!important}#outlook a{padding:0}img{-ms-interpolation-mode:bicubic}table{mso-table-lspace:0;mso-table-rspace:0}.ReadMsgBody{width:100%}.ExternalClass{width:100%}p,a,li,td,blockquote{mso-line-height-rule:exactly}a[href^=tel],a[href^=sms]{color:inherit;cursor:default;text-decoration:none}p,a,li,td,body,table,blockquote{-ms-text-size-adjust:100%;-webkit-text-size-adjust:100%}.ExternalClass,.ExternalClass p,.ExternalClass td,.ExternalClass div,.ExternalClass span,.ExternalClass font{line-height:100%}a[x-apple-data-detectors]{color:inherit!important;text-decoration:none!important;font-size:inherit!important;font-family:inherit!important;font-weight:inherit!important;line-height:inherit!important}#bodyCell{padding:10px}.templateContainer{max-width:600px!important}a.mcnButton{display:block}.mcnImage,.mcnRetinaImage{vertical-align:bottom}.mcnTextContent{word-break:break-word}.mcnTextContent img{height:auto!important}.mcnDividerBlock{table-layout:fixed!important}body,#bodyTable{background-color:#fafafa}#bodyCell{border-top:0}.templateContainer{border:0}h1{color:#202020;font-family:Helvetica;font-size:26px;font-style:normal;font-weight:bold;line-height:125%;letter-spacing:normal;text-align:left}h2{color:#202020;font-family:Helvetica;font-size:22px;font-style:normal;font-weight:bold;line-height:125%;letter-spacing:normal;text-align:left}h3{color:#202020;font-family:Helvetica;font-size:20px;font-style:normal;font-weight:bold;line-height:125%;letter-spacing:normal;text-align:left}h4{color:#202020;font-family:Helvetica;font-size:18px;font-style:normal;font-weight:bold;line-height:125%;letter-spacing:normal;text-align:left}#templatePreheader{background-color:#843bae;background-image:none;background-repeat:no-repeat;background-position:center;background-size:cover;border-top:0;border-bottom:0;padding-top:2px;padding-bottom:0}#templatePreheader .mcnTextContent,#templatePreheader .mcnTextContent p{color:#656565;font-family:Helvetica;font-size:0;line-height:150%;text-align:center}#templatePreheader .mcnTextContent a,#templatePreheader .mcnTextContent p a{color:#fff;font-weight:normal;text-decoration:none}#templateHeader{background-color:#fff;background-image:none;background-repeat:no-repeat;background-position:center;background-size:cover;border-top:0;border-bottom:0;padding-top:9px;padding-bottom:0}#templateHeader .mcnTextContent,#templateHeader .mcnTextContent p{color:#202020;font-family:Helvetica;font-size:16px;line-height:150%;text-align:left}#templateHeader .mcnTextContent a,#templateHeader .mcnTextContent p a{color:#007c89;font-weight:normal;text-decoration:underline}#templateBody{background-color:#fff;background-image:none;background-repeat:no-repeat;background-position:center;background-size:cover;border-top:0;border-bottom:2px solid #eaeaea;padding-top:0;padding-bottom:9px}#templateBody .mcnTextContent,#templateBody .mcnTextContent p{color:#202020;font-family:Helvetica;font-size:16px;line-height:150%;text-align:left}#templateBody .mcnTextContent a,#templateBody .mcnTextContent p a{color:#007c89;font-weight:normal;text-decoration:underline}#templateFooter{background-color:#fafafa;background-image:none;background-repeat:no-repeat;background-position:center;background-size:cover;border-top:0;border-bottom:0;padding-top:9px;padding-bottom:9px}#templateFooter .mcnTextContent,#templateFooter .mcnTextContent p{color:#656565;font-family:Helvetica;font-size:12px;line-height:150%;text-align:center}#templateFooter .mcnTextContent a,#templateFooter .mcnTextContent p a{color:#656565;font-weight:normal;text-decoration:underline}@media only screen and (min-width:768px){.templateContainer{width:600px!important}}@media only screen and (max-width:480px){body,table,td,p,a,li,blockquote{-webkit-text-size-adjust:none!important}}@media only screen and (max-width:480px){body{width:100%!important;min-width:100%!important}}@media only screen and (max-width:480px){.mcnRetinaImage{max-width:100%!important}}@media only screen and (max-width:480px){.mcnImage{width:100%!important}}@media only screen and (max-width:480px){.mcnCartContainer,.mcnCaptionTopContent,.mcnRecContentContainer,.mcnCaptionBottomContent,.mcnTextContentContainer,.mcnBoxedTextContentContainer,.mcnImageGroupContentContainer,.mcnCaptionLeftTextContentContainer,.mcnCaptionRightTextContentContainer,.mcnCaptionLeftImageContentContainer,.mcnCaptionRightImageContentContainer,.mcnImageCardLeftTextContentContainer,.mcnImageCardRightTextContentContainer,.mcnImageCardLeftImageContentContainer,.mcnImageCardRightImageContentContainer{max-width:100%!important;width:100%!important}}@media only screen and (max-width:480px){.mcnBoxedTextContentContainer{min-width:100%!important}}@media only screen and (max-width:480px){.mcnImageGroupContent{padding:9px!important}}@media only screen and (max-width:480px){.mcnCaptionLeftContentOuter .mcnTextContent,.mcnCaptionRightContentOuter .mcnTextContent{padding-top:9px!important}}@media only screen and (max-width:480px){.mcnImageCardTopImageContent,.mcnCaptionBottomContent:last-child .mcnCaptionBottomImageContent,.mcnCaptionBlockInner .mcnCaptionTopContent:last-child .mcnTextContent{padding-top:18px!important}}@media only screen and (max-width:480px){.mcnImageCardBottomImageContent{padding-bottom:9px!important}}@media only screen and (max-width:480px){.mcnImageGroupBlockInner{padding-top:0!important;padding-bottom:0!important}}@media only screen and (max-width:480px){.mcnImageGroupBlockOuter{padding-top:9px!important;padding-bottom:9px!important}}@media only screen and (max-width:480px){.mcnTextContent,.mcnBoxedTextContentColumn{padding-right:18px!important;padding-left:18px!important}}@media only screen and (max-width:480px){.mcnImageCardLeftImageContent,.mcnImageCardRightImageContent{padding-right:18px!important;padding-bottom:0!important;padding-left:18px!important}}@media only screen and (max-width:480px){.mcpreview-image-uploader{display:none!important;width:100%!important}}@media only screen and (max-width:480px){h1{font-size:22px!important;line-height:125%!important}}@media only screen and (max-width:480px){h2{font-size:20px!important;line-height:125%!important}}@media only screen and (max-width:480px){h3{font-size:18px!important;line-height:125%!important}}@media only screen and (max-width:480px){h4{font-size:16px!important;line-height:150%!important}}@media only screen and (max-width:480px){.mcnBoxedTextContentContainer .mcnTextContent,.mcnBoxedTextContentContainer .mcnTextContent p{font-size:14px!important;line-height:150%!important}}@media only screen and (max-width:480px){#templatePreheader{display:block!important}}@media only screen and (max-width:480px){#templatePreheader .mcnTextContent,#templatePreheader .mcnTextContent p{font-size:14px!important;line-height:150%!important}}@media only screen and (max-width:480px){#templateHeader .mcnTextContent,#templateHeader .mcnTextContent p{font-size:16px!important;line-height:150%!important}}@media only screen and (max-width:480px){#templateBody .mcnTextContent,#templateBody .mcnTextContent p{font-size:16px!important;line-height:150%!important}}@media only screen and (max-width:480px){#templateFooter .mcnTextContent,#templateFooter .mcnTextContent p{font-size:14px!important;line-height:150%!important}}</style></head>
	<body>
	<!--[if !gte mso 9]><span class=mcnPreviewText style=display:none;font-size:0;line-height:0;max-height:0;max-width:0;opacity:0;overflow:hidden;visibility:hidden;mso-hide:all>*|MC_PREVIEW_TEXT|*</span><!--<![endif]-->
	<center>
	<table align=center border=0 cellpadding=0 cellspacing=0 height=100% width=100% id=bodyTable>
	<tr>
	<td align=center valign=top id=bodyCell>
	<!--[if (gte mso 9)|(IE)]>
	<table align=center border=0 cellspacing=0 cellpadding=0 width=600 style=width:600px>
	<tr>
	<td align=center valign=top width=600 style=width:600px>
	<![endif]-->
	<table border=0 cellpadding=0 cellspacing=0 width=100% class=templateContainer>
	<tr>
	<td valign=top id=templatePreheader><table border=0 cellpadding=0 cellspacing=0 width=100% class=mcnTextBlock style=min-width:100%>
	<tbody class=mcnTextBlockOuter>
	<tr>
	<td valign=top class=mcnTextBlockInner>
	<!--[if mso]>
	<table align=left border=0 cellspacing=0 cellpadding=0 width=100% style=width:100%>
	<tr>
	<![endif]-->
	<!--[if mso]>
	<td valign=top width=600 style=width:600px>
	<![endif]-->
	<table align=left border=0 cellpadding=0 cellspacing=0 style=max-width:100%;min-width:100% width=100% class=mcnTextContentContainer>
	<tbody><tr>
	<td valign=top class=mcnTextContent style="padding:0 18px 5px;text-align:center">
	</td>
	</tr>
	</tbody></table>
	<!--[if mso]>
	</td>
	<![endif]-->
	<!--[if mso]>
	</tr>
	</table>
	<![endif]-->
	</td>
	</tr>
	</tbody>
	</table></td>
	</tr>
	<tr>
	<td valign=top id=templateHeader><table border=0 cellpadding=0 cellspacing=0 width=100% class=mcnImageBlock style=min-width:100%>
	<tbody class=mcnImageBlockOuter>
	<tr>
	<td valign=top style=padding:9px class=mcnImageBlockInner>
	<table align=left width=100% border=0 cellpadding=0 cellspacing=0 class=mcnImageContentContainer style=min-width:100%>
	<tbody><tr>
	<td class=mcnImageContent valign=top style=padding-right:9px;padding-left:9px;padding-top:0;padding-bottom:0;text-align:center>
	<img align=center alt src=https://mcusercontent.com/b5abe213da603aaf115c53642/images/8ad6b1dd-1f21-449a-b209-907dcc4a11c0.png width=100 style="max-width:100px;padding-bottom:0;vertical-align:bottom;display:inline!important;border:10px none #f00" class=mcnImage>
	</td>
	</tr>
	</tbody></table>
	</td>
	</tr>
	</tbody>
	</table></td>
	</tr>
	<tr>
	<td valign=top id=templateBody><table border=0 cellpadding=0 cellspacing=0 width=100% class=mcnDividerBlock style=min-width:100%>
	<tbody class=mcnDividerBlockOuter>
	<tr>
	<td class=mcnDividerBlockInner style=min-width:100%;padding:18px>
	<table class=mcnDividerContent border=0 cellpadding=0 cellspacing=0 width=100% style="min-width:100%;border-top:2px solid #eaeaea">
	<tbody><tr>
	<td>
	<span></span>
	</td>
	</tr>
	</tbody></table>
	</td>
	</tr>
	</tbody>
	</table><table border=0 cellpadding=0 cellspacing=0 width=100% class=mcnTextBlock style=min-width:100%>
	<tbody class=mcnTextBlockOuter>
	<tr>
	<td valign=top class=mcnTextBlockInner style=padding-top:9px>
	<!--[if mso]>
	<table align=left border=0 cellspacing=0 cellpadding=0 width=100% style=width:100%>
	<tr>
	<![endif]-->
	<!--[if mso]>
	<td valign=top width=600 style=width:600px>
	<![endif]-->
	<table align=left border=0 cellpadding=0 cellspacing=0 style=max-width:100%;min-width:100% width=100% class=mcnTextContentContainer>
	<tbody><tr>
	<td valign=top class=mcnTextContent style=padding-top:0;padding-right:18px;padding-bottom:9px;padding-left:18px>
	<h1>User Registration<br>
	&nbsp;</h1>
	<p>Dear&nbsp;{{.name}},<br>
	Selamat registrasi anda menggunakan email telah berhasil<br>
	silahkan klik link dibawah ini untuk melakukan verifikasi email&nbsp;<br>
	<br>
	{{.link}}<br>
	&nbsp;</p>
	<p><span style=font-size:10px>Jangan membalas ke email ini. Email ini hanya diperuntukan sebagai informasi saja.<br>
	Ruang Nyaman</span></p>
	</td>
	</tr>
	</tbody></table>
	<!--[if mso]>
	</td>
	<![endif]-->
	<!--[if mso]>
	</tr>
	</table>
	<![endif]-->
	</td>
	</tr>
	</tbody>
	</table><table border=0 cellpadding=0 cellspacing=0 width=100% class=mcnDividerBlock style=min-width:100%>
	<tbody class=mcnDividerBlockOuter>
	<tr>
	<td class=mcnDividerBlockInner style=min-width:100%;padding:18px>
	<table class=mcnDividerContent border=0 cellpadding=0 cellspacing=0 width=100% style="min-width:100%;border-top:2px solid #eaeaea">
	<tbody><tr>
	<td>
	<span></span>
	</td>
	</tr>
	</tbody></table>
	</td>
	</tr>
	</tbody>
	</table><table border=0 cellpadding=0 cellspacing=0 width=100% class=mcnTextBlock style=min-width:100%>
	<tbody class=mcnTextBlockOuter>
	<tr>
	<td valign=top class=mcnTextBlockInner style=padding-top:9px>
	<!--[if mso]>
	<table align=left border=0 cellspacing=0 cellpadding=0 width=100% style=width:100%>
	<tr>
	<![endif]-->
	<!--[if mso]>
	<td valign=top width=600 style=width:600px>
	<![endif]-->
	<table align=left border=0 cellpadding=0 cellspacing=0 style=max-width:100%;min-width:100% width=100% class=mcnTextContentContainer>
	<tbody><tr>
	<td valign=top class=mcnTextContent style=padding-top:0;padding-right:18px;padding-bottom:9px;padding-left:18px>
	<h1>User Registrasi<br>
	&nbsp;</h1>
	<p>Dear {{.name}},<br>
	congratulation, your registration using email succesful!<br>
	please clik link below for email verification&nbsp;<br>
	<br>
	{{.link}}<br>
	&nbsp;</p>
	<p><span style=font-size:10px>Please do not reply to this email. This email only intended as information only.<br>
	Ruang Nyaman</span></p>
	</td>
	</tr>
	</tbody></table>
	<!--[if mso]>
	</td>
	<![endif]-->
	<!--[if mso]>
	</tr>
	</table>
	<![endif]-->
	</td>
	</tr>
	</tbody>
	</table><table border=0 cellpadding=0 cellspacing=0 width=100% class=mcnBoxedTextBlock style=min-width:100%>
	<!--[if gte mso 9]>
	<table align=center border=0 cellspacing=0 cellpadding=0 width=100%>
	<![endif]-->
	<tbody class=mcnBoxedTextBlockOuter>
	<tr>
	<td valign=top class=mcnBoxedTextBlockInner>
	<!--[if gte mso 9]>
	<td align=center valign=top ">
	<![endif]-->
	<table align=left border=0 cellpadding=0 cellspacing=0 width=100% style=min-width:100% class=mcnBoxedTextContentContainer>
	<tbody><tr>
	<td style=padding-top:9px;padding-left:18px;padding-bottom:9px;padding-right:18px>
	<table border=0 cellspacing=0 class=mcnTextContentContainer width=100% style="min-width:100%!important;background-color:#843bae;border:1px solid">
	<tbody><tr>
	<td valign=top class=mcnTextContent style=padding:18px;color:#f2f2f2;font-family:Helvetica;font-size:14px;font-weight:normal;text-align:center>
	<em>Copyright © 2021 Ruang Nyaman , All rights reserved.</em>
	</td>
	</tr>
	</tbody></table>
	</td>
	</tr>
	</tbody></table>
	<!--[if gte mso 9]>
	</td>
	<![endif]-->
	<!--[if gte mso 9]>
	</tr>
	</table>
	<![endif]-->
	</td>
	</tr>
	</tbody>
	</table></td>
	</tr>
	<tr>
	<td valign=top id=templateFooter><table border=0 cellpadding=0 cellspacing=0 width=100% class=mcnDividerBlock style=min-width:100%>
	<tbody class=mcnDividerBlockOuter>
	<tr>
	<td class=mcnDividerBlockInner style="min-width:100%;padding:10px 18px 25px">
	<table class=mcnDividerContent border=0 cellpadding=0 cellspacing=0 width=100% style="min-width:100%;border-top:2px solid #eee">
	<tbody><tr>
	<td>
	<span></span>
	</td>
	</tr>
	</tbody></table>
	</td>
	</tr>
	</tbody>
	</table></td>
	</tr>
	</table>
	<!--[if (gte mso 9)|(IE)]>
	</td>
	</tr>
	</table>
	<![endif]-->
	</td>
	</tr>
	</table>
	</center>
	</body>
	</html>`

	TemplateLostPass = `
	<!doctype html>
<html xmlns=http://www.w3.org/1999/xhtml xmlns:v=urn:schemas-microsoft-com:vml xmlns:o=urn:schemas-microsoft-com:office:office>
<head>
<!--[if gte mso 15]>
<xml>
<o:OfficeDocumentSettings>
<o:AllowPNG/>
<o:PixelsPerInch>96</o:PixelsPerInch>
</o:OfficeDocumentSettings>
</xml>
<![endif]-->
<meta charset=UTF-8>
<meta http-equiv=X-UA-Compatible content="IE=edge">
<meta name=viewport content="width=device-width, initial-scale=1">
<title>*|MC:SUBJECT|*</title>
<style type=text/css>p{margin:10px 0;padding:0}table{border-collapse:collapse}h1,h2,h3,h4,h5,h6{display:block;margin:0;padding:0}img,a img{border:0;height:auto;outline:0;text-decoration:none}body,#bodyTable,#bodyCell{height:100%;margin:0;padding:0;width:100%}.mcnPreviewText{display:none!important}#outlook a{padding:0}img{-ms-interpolation-mode:bicubic}table{mso-table-lspace:0;mso-table-rspace:0}.ReadMsgBody{width:100%}.ExternalClass{width:100%}p,a,li,td,blockquote{mso-line-height-rule:exactly}a[href^=tel],a[href^=sms]{color:inherit;cursor:default;text-decoration:none}p,a,li,td,body,table,blockquote{-ms-text-size-adjust:100%;-webkit-text-size-adjust:100%}.ExternalClass,.ExternalClass p,.ExternalClass td,.ExternalClass div,.ExternalClass span,.ExternalClass font{line-height:100%}a[x-apple-data-detectors]{color:inherit!important;text-decoration:none!important;font-size:inherit!important;font-family:inherit!important;font-weight:inherit!important;line-height:inherit!important}#bodyCell{padding:10px}.templateContainer{max-width:600px!important}a.mcnButton{display:block}.mcnImage,.mcnRetinaImage{vertical-align:bottom}.mcnTextContent{word-break:break-word}.mcnTextContent img{height:auto!important}.mcnDividerBlock{table-layout:fixed!important}body,#bodyTable{background-color:#fafafa}#bodyCell{border-top:0}.templateContainer{border:0}h1{color:#202020;font-family:Helvetica;font-size:26px;font-style:normal;font-weight:bold;line-height:125%;letter-spacing:normal;text-align:left}h2{color:#202020;font-family:Helvetica;font-size:22px;font-style:normal;font-weight:bold;line-height:125%;letter-spacing:normal;text-align:left}h3{color:#202020;font-family:Helvetica;font-size:20px;font-style:normal;font-weight:bold;line-height:125%;letter-spacing:normal;text-align:left}h4{color:#202020;font-family:Helvetica;font-size:18px;font-style:normal;font-weight:bold;line-height:125%;letter-spacing:normal;text-align:left}#templatePreheader{background-color:#843bae;background-image:none;background-repeat:no-repeat;background-position:center;background-size:cover;border-top:0;border-bottom:0;padding-top:2px;padding-bottom:0}#templatePreheader .mcnTextContent,#templatePreheader .mcnTextContent p{color:#656565;font-family:Helvetica;font-size:0;line-height:150%;text-align:center}#templatePreheader .mcnTextContent a,#templatePreheader .mcnTextContent p a{color:#fff;font-weight:normal;text-decoration:none}#templateHeader{background-color:#fff;background-image:none;background-repeat:no-repeat;background-position:center;background-size:cover;border-top:0;border-bottom:0;padding-top:9px;padding-bottom:0}#templateHeader .mcnTextContent,#templateHeader .mcnTextContent p{color:#202020;font-family:Helvetica;font-size:16px;line-height:150%;text-align:left}#templateHeader .mcnTextContent a,#templateHeader .mcnTextContent p a{color:#007c89;font-weight:normal;text-decoration:underline}#templateBody{background-color:#fff;background-image:none;background-repeat:no-repeat;background-position:center;background-size:cover;border-top:0;border-bottom:2px solid #eaeaea;padding-top:0;padding-bottom:9px}#templateBody .mcnTextContent,#templateBody .mcnTextContent p{color:#202020;font-family:Helvetica;font-size:16px;line-height:150%;text-align:left}#templateBody .mcnTextContent a,#templateBody .mcnTextContent p a{color:#007c89;font-weight:normal;text-decoration:underline}#templateFooter{background-color:#fafafa;background-image:none;background-repeat:no-repeat;background-position:center;background-size:cover;border-top:0;border-bottom:0;padding-top:9px;padding-bottom:9px}#templateFooter .mcnTextContent,#templateFooter .mcnTextContent p{color:#656565;font-family:Helvetica;font-size:12px;line-height:150%;text-align:center}#templateFooter .mcnTextContent a,#templateFooter .mcnTextContent p a{color:#656565;font-weight:normal;text-decoration:underline}@media only screen and (min-width:768px){.templateContainer{width:600px!important}}@media only screen and (max-width:480px){body,table,td,p,a,li,blockquote{-webkit-text-size-adjust:none!important}}@media only screen and (max-width:480px){body{width:100%!important;min-width:100%!important}}@media only screen and (max-width:480px){.mcnRetinaImage{max-width:100%!important}}@media only screen and (max-width:480px){.mcnImage{width:100%!important}}@media only screen and (max-width:480px){.mcnCartContainer,.mcnCaptionTopContent,.mcnRecContentContainer,.mcnCaptionBottomContent,.mcnTextContentContainer,.mcnBoxedTextContentContainer,.mcnImageGroupContentContainer,.mcnCaptionLeftTextContentContainer,.mcnCaptionRightTextContentContainer,.mcnCaptionLeftImageContentContainer,.mcnCaptionRightImageContentContainer,.mcnImageCardLeftTextContentContainer,.mcnImageCardRightTextContentContainer,.mcnImageCardLeftImageContentContainer,.mcnImageCardRightImageContentContainer{max-width:100%!important;width:100%!important}}@media only screen and (max-width:480px){.mcnBoxedTextContentContainer{min-width:100%!important}}@media only screen and (max-width:480px){.mcnImageGroupContent{padding:9px!important}}@media only screen and (max-width:480px){.mcnCaptionLeftContentOuter .mcnTextContent,.mcnCaptionRightContentOuter .mcnTextContent{padding-top:9px!important}}@media only screen and (max-width:480px){.mcnImageCardTopImageContent,.mcnCaptionBottomContent:last-child .mcnCaptionBottomImageContent,.mcnCaptionBlockInner .mcnCaptionTopContent:last-child .mcnTextContent{padding-top:18px!important}}@media only screen and (max-width:480px){.mcnImageCardBottomImageContent{padding-bottom:9px!important}}@media only screen and (max-width:480px){.mcnImageGroupBlockInner{padding-top:0!important;padding-bottom:0!important}}@media only screen and (max-width:480px){.mcnImageGroupBlockOuter{padding-top:9px!important;padding-bottom:9px!important}}@media only screen and (max-width:480px){.mcnTextContent,.mcnBoxedTextContentColumn{padding-right:18px!important;padding-left:18px!important}}@media only screen and (max-width:480px){.mcnImageCardLeftImageContent,.mcnImageCardRightImageContent{padding-right:18px!important;padding-bottom:0!important;padding-left:18px!important}}@media only screen and (max-width:480px){.mcpreview-image-uploader{display:none!important;width:100%!important}}@media only screen and (max-width:480px){h1{font-size:22px!important;line-height:125%!important}}@media only screen and (max-width:480px){h2{font-size:20px!important;line-height:125%!important}}@media only screen and (max-width:480px){h3{font-size:18px!important;line-height:125%!important}}@media only screen and (max-width:480px){h4{font-size:16px!important;line-height:150%!important}}@media only screen and (max-width:480px){.mcnBoxedTextContentContainer .mcnTextContent,.mcnBoxedTextContentContainer .mcnTextContent p{font-size:14px!important;line-height:150%!important}}@media only screen and (max-width:480px){#templatePreheader{display:block!important}}@media only screen and (max-width:480px){#templatePreheader .mcnTextContent,#templatePreheader .mcnTextContent p{font-size:14px!important;line-height:150%!important}}@media only screen and (max-width:480px){#templateHeader .mcnTextContent,#templateHeader .mcnTextContent p{font-size:16px!important;line-height:150%!important}}@media only screen and (max-width:480px){#templateBody .mcnTextContent,#templateBody .mcnTextContent p{font-size:16px!important;line-height:150%!important}}@media only screen and (max-width:480px){#templateFooter .mcnTextContent,#templateFooter .mcnTextContent p{font-size:14px!important;line-height:150%!important}}</style></head>
<body>
<!--[if !gte mso 9]><span class=mcnPreviewText style=display:none;font-size:0;line-height:0;max-height:0;max-width:0;opacity:0;overflow:hidden;visibility:hidden;mso-hide:all>*|MC_PREVIEW_TEXT|*</span><!--<![endif]-->
<center>
<table align=center border=0 cellpadding=0 cellspacing=0 height=100% width=100% id=bodyTable>
<tr>
<td align=center valign=top id=bodyCell>
<!--[if (gte mso 9)|(IE)]>
<table align=center border=0 cellspacing=0 cellpadding=0 width=600 style=width:600px>
<tr>
<td align=center valign=top width=600 style=width:600px>
<![endif]-->
<table border=0 cellpadding=0 cellspacing=0 width=100% class=templateContainer>
<tr>
<td valign=top id=templatePreheader><table border=0 cellpadding=0 cellspacing=0 width=100% class=mcnTextBlock style=min-width:100%>
<tbody class=mcnTextBlockOuter>
<tr>
<td valign=top class=mcnTextBlockInner style=padding-top:9px>
<!--[if mso]>
<table align=left border=0 cellspacing=0 cellpadding=0 width=100% style=width:100%>
<tr>
<![endif]-->
<!--[if mso]>
<td valign=top width=600 style=width:600px>
<![endif]-->
<table align=left border=0 cellpadding=0 cellspacing=0 style=max-width:100%;min-width:100% width=100% class=mcnTextContentContainer>
<tbody><tr>
<td valign=top class=mcnTextContent style="padding:0 18px 9px;text-align:center">
</td>
</tr>
</tbody></table>
<!--[if mso]>
</td>
<![endif]-->
<!--[if mso]>
</tr>
</table>
<![endif]-->
</td>
</tr>
</tbody>
</table></td>
</tr>
<tr>
<td valign=top id=templateHeader><table border=0 cellpadding=0 cellspacing=0 width=100% class=mcnImageBlock style=min-width:100%>
<tbody class=mcnImageBlockOuter>
<tr>
<td valign=top style=padding:9px class=mcnImageBlockInner>
<table align=left width=100% border=0 cellpadding=0 cellspacing=0 class=mcnImageContentContainer style=min-width:100%>
<tbody><tr>
<td class=mcnImageContent valign=top style=padding-right:9px;padding-left:9px;padding-top:0;padding-bottom:0;text-align:center>
<img align=center alt src=https://mcusercontent.com/b5abe213da603aaf115c53642/images/8ad6b1dd-1f21-449a-b209-907dcc4a11c0.png width=100 style="max-width:100px;padding-bottom:0;vertical-align:bottom;display:inline!important;border:10px none #f00" class=mcnImage>
</td>
</tr>
</tbody></table>
</td>
</tr>
</tbody>
</table></td>
</tr>
<tr>
<td valign=top id=templateBody><table border=0 cellpadding=0 cellspacing=0 width=100% class=mcnDividerBlock style=min-width:100%>
<tbody class=mcnDividerBlockOuter>
<tr>
<td class=mcnDividerBlockInner style=min-width:100%;padding:18px>
<table class=mcnDividerContent border=0 cellpadding=0 cellspacing=0 width=100% style="min-width:100%;border-top:2px solid #eaeaea">
<tbody><tr>
<td>
<span></span>
</td>
</tr>
</tbody></table>
</td>
</tr>
</tbody>
</table><table border=0 cellpadding=0 cellspacing=0 width=100% class=mcnTextBlock style=min-width:100%>
<tbody class=mcnTextBlockOuter>
<tr>
<td valign=top class=mcnTextBlockInner style=padding-top:9px>
<!--[if mso]>
<table align=left border=0 cellspacing=0 cellpadding=0 width=100% style=width:100%>
<tr>
<![endif]-->
<!--[if mso]>
<td valign=top width=600 style=width:600px>
<![endif]-->
<table align=left border=0 cellpadding=0 cellspacing=0 style=max-width:100%;min-width:100% width=100% class=mcnTextContentContainer>
<tbody><tr>
<td valign=top class=mcnTextContent style=padding-top:0;padding-right:18px;padding-bottom:9px;padding-left:18px>
<h1>Reset Password<br>
&nbsp;</h1>
<p>Kami menerima permintaan Reset Password dari akun anda, silahkan gunakan kode verifikasi dibawah :<br>
<br>
<strong>{{.code}}</strong><br>
<br>
Kode ini hanya valid selama 15 menit, pastikan anda menyelesaikan semua proses di bawah 15 menit.</p>
<p><br>
Mohon tidak membagikan kode ini dengan siapapun. Jika anda tidak melakukan permintaan ini, maka abaikan email ini atau hubungi support&nbsp;ruang nyaman.</p>
<p>&nbsp;</p>
<p><span style=font-size:10px>Jangan membalas ke email ini. Email ini hanya diperuntukan sebagai informasi saja.<br>
Ruang Nyaman</span></p>
</td>
</tr>
</tbody></table>
<!--[if mso]>
</td>
<![endif]-->
<!--[if mso]>
</tr>
</table>
<![endif]-->
</td>
</tr>
</tbody>
</table><table border=0 cellpadding=0 cellspacing=0 width=100% class=mcnDividerBlock style=min-width:100%>
<tbody class=mcnDividerBlockOuter>
<tr>
<td class=mcnDividerBlockInner style=min-width:100%;padding:18px>
<table class=mcnDividerContent border=0 cellpadding=0 cellspacing=0 width=100% style="min-width:100%;border-top:2px solid #eaeaea">
<tbody><tr>
<td>
<span></span>
</td>
</tr>
</tbody></table>
</td>
</tr>
</tbody>
</table><table border=0 cellpadding=0 cellspacing=0 width=100% class=mcnTextBlock style=min-width:100%>
<tbody class=mcnTextBlockOuter>
<tr>
<td valign=top class=mcnTextBlockInner style=padding-top:9px>
<!--[if mso]>
<table align=left border=0 cellspacing=0 cellpadding=0 width=100% style=width:100%>
<tr>
<![endif]-->
<!--[if mso]>
<td valign=top width=600 style=width:600px>
<![endif]-->
<table align=left border=0 cellpadding=0 cellspacing=0 style=max-width:100%;min-width:100% width=100% class=mcnTextContentContainer>
<tbody><tr>
<td valign=top class=mcnTextContent style=padding-top:0;padding-right:18px;padding-bottom:9px;padding-left:18px>
<h1>Reset Password<br>
&nbsp;</h1>
<p>We received request for&nbsp;Reset Password from your account, please use this verification code below :<br>
<br>
<strong>{{.code}}</strong><br>
<br>
this code only valid for 15 minutes, make sure you complete the change procedure in less than 15 minutes.</p>
<p>please do not share this code with others. if this request wasn't from you, please ignore this email or contact ruang nyaman support</p>
<p>&nbsp;</p>
<p><span style=font-size:10px>Please do not reply to this email. This email only intended as information only.<br>
Ruang Nyaman</span></p>
</td>
</tr>
</tbody></table>
<!--[if mso]>
</td>
<![endif]-->
<!--[if mso]>
</tr>
</table>
<![endif]-->
</td>
</tr>
</tbody>
</table><table border=0 cellpadding=0 cellspacing=0 width=100% class=mcnBoxedTextBlock style=min-width:100%>
<!--[if gte mso 9]>
<table align=center border=0 cellspacing=0 cellpadding=0 width=100%>
<![endif]-->
<tbody class=mcnBoxedTextBlockOuter>
<tr>
<td valign=top class=mcnBoxedTextBlockInner>
<!--[if gte mso 9]>
<td align=center valign=top ">
<![endif]-->
<table align=left border=0 cellpadding=0 cellspacing=0 width=100% style=min-width:100% class=mcnBoxedTextContentContainer>
<tbody><tr>
<td style=padding-top:9px;padding-left:18px;padding-bottom:9px;padding-right:18px>
<table border=0 cellspacing=0 class=mcnTextContentContainer width=100% style="min-width:100%!important;background-color:#843bae;border:1px solid">
<tbody><tr>
<td valign=top class=mcnTextContent style=padding:18px;color:#f2f2f2;font-family:Helvetica;font-size:14px;font-weight:normal;text-align:center>
<em>Copyright © 2021 Ruang Nyaman , All rights reserved.</em>
</td>
</tr>
</tbody></table>
</td>
</tr>
</tbody></table>
<!--[if gte mso 9]>
</td>
<![endif]-->
<!--[if gte mso 9]>
</tr>
</table>
<![endif]-->
</td>
</tr>
</tbody>
</table></td>
</tr>
<tr>
<td valign=top id=templateFooter><table border=0 cellpadding=0 cellspacing=0 width=100% class=mcnDividerBlock style=min-width:100%>
<tbody class=mcnDividerBlockOuter>
<tr>
<td class=mcnDividerBlockInner style="min-width:100%;padding:10px 18px 25px">
<table class=mcnDividerContent border=0 cellpadding=0 cellspacing=0 width=100% style="min-width:100%;border-top:2px solid #eee">
<tbody><tr>
<td>
<span></span>
</td>
</tr>
</tbody></table>
</td>
</tr>
</tbody>
</table></td>
</tr>
</table>
<!--[if (gte mso 9)|(IE)]>
</td>
</tr>
</table>
<![endif]-->
</td>
</tr>
</table>
</center>
</body>
</html>`
)
