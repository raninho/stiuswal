package fixture

const HTMLCrawler = ` 













<html>





























































<head>
<meta http-equiv="Content-Type" content="text/html; charset=UTF-8">

	








	<script language="javascript" type="text/JavaScript" src="https://www2.tjal.jus.br/sajcas/verificarLogin.js?script=1564434719419"></script>

	<script language="javascript">
		if(window.sajcas && window.sajcas.usuarioLogadoNoCasServer) {
			var urlRetornoSistema = '';
			if(urlRetornoSistema === '') {
				urlRetornoSistema = window.location.href;
			}

			if(urlRetornoSistema.lastIndexOf('?') != -1) {
				urlRetornoSistema += '&';
			} else {
				urlRetornoSistema += '?';
			}
			urlRetornoSistema+= 'gateway=true';
			window.location.href = urlRetornoSistema;
		}
	</script>


	<title>Portal de Serviços e-SAJ</title>
	































<!-- Includes do spw -->

<script type="text/javascript">window.NREUM||(NREUM={}),__nr_require=function(e,n,t){function r(t){if(!n[t]){var o=n[t]={exports:{}};e[t][0].call(o.exports,function(n){var o=e[t][1][n];return r(o||n)},o,o.exports)}return n[t].exports}if("function"==typeof __nr_require)return __nr_require;for(var o=0;o<t.length;o++)r(t[o]);return r}({1:[function(e,n,t){function r(){}function o(e,n,t){return function(){return i(e,[c.now()].concat(u(arguments)),n?null:this,t),n?void 0:this}}var i=e("handle"),a=e(3),u=e(4),f=e("ee").get("tracer"),c=e("loader"),s=NREUM;"undefined"==typeof window.newrelic&&(newrelic=s);var p=["setPageViewName","setCustomAttribute","setErrorHandler","finished","addToTrace","inlineHit","addRelease"],d="api-",l=d+"ixn-";a(p,function(e,n){s[n]=o(d+n,!0,"api")}),s.addPageAction=o(d+"addPageAction",!0),s.setCurrentRouteName=o(d+"routeName",!0),n.exports=newrelic,s.interaction=function(){return(new r).get()};var m=r.prototype={createTracer:function(e,n){var t={},r=this,o="function"==typeof n;return i(l+"tracer",[c.now(),e,t],r),function(){if(f.emit((o?"":"no-")+"fn-start",[c.now(),r,o],t),o)try{return n.apply(this,arguments)}catch(e){throw f.emit("fn-err",[arguments,this,e],t),e}finally{f.emit("fn-end",[c.now()],t)}}}};a("actionText,setName,setAttribute,save,ignore,onEnd,getContext,end,get".split(","),function(e,n){m[n]=o(l+n)}),newrelic.noticeError=function(e,n){"string"==typeof e&&(e=new Error(e)),i("err",[e,c.now(),!1,n])}},{}],2:[function(e,n,t){function r(e,n){if(!o)return!1;if(e!==o)return!1;if(!n)return!0;if(!i)return!1;for(var t=i.split("."),r=n.split("."),a=0;a<r.length;a++)if(r[a]!==t[a])return!1;return!0}var o=null,i=null,a=/Version\/(\S+)\s+Safari/;if(navigator.userAgent){var u=navigator.userAgent,f=u.match(a);f&&u.indexOf("Chrome")===-1&&u.indexOf("Chromium")===-1&&(o="Safari",i=f[1])}n.exports={agent:o,version:i,match:r}},{}],3:[function(e,n,t){function r(e,n){var t=[],r="",i=0;for(r in e)o.call(e,r)&&(t[i]=n(r,e[r]),i+=1);return t}var o=Object.prototype.hasOwnProperty;n.exports=r},{}],4:[function(e,n,t){function r(e,n,t){n||(n=0),"undefined"==typeof t&&(t=e?e.length:0);for(var r=-1,o=t-n||0,i=Array(o<0?0:o);++r<o;)i[r]=e[n+r];return i}n.exports=r},{}],5:[function(e,n,t){n.exports={exists:"undefined"!=typeof window.performance&&window.performance.timing&&"undefined"!=typeof window.performance.timing.navigationStart}},{}],ee:[function(e,n,t){function r(){}function o(e){function n(e){return e&&e instanceof r?e:e?f(e,u,i):i()}function t(t,r,o,i){if(!d.aborted||i){e&&e(t,r,o);for(var a=n(o),u=v(t),f=u.length,c=0;c<f;c++)u[c].apply(a,r);var p=s[y[t]];return p&&p.push([b,t,r,a]),a}}function l(e,n){h[e]=v(e).concat(n)}function m(e,n){var t=h[e];if(t)for(var r=0;r<t.length;r++)t[r]===n&&t.splice(r,1)}function v(e){return h[e]||[]}function g(e){return p[e]=p[e]||o(t)}function w(e,n){c(e,function(e,t){n=n||"feature",y[t]=n,n in s||(s[n]=[])})}var h={},y={},b={on:l,addEventListener:l,removeEventListener:m,emit:t,get:g,listeners:v,context:n,buffer:w,abort:a,aborted:!1};return b}function i(){return new r}function a(){(s.api||s.feature)&&(d.aborted=!0,s=d.backlog={})}var u="nr@context",f=e("gos"),c=e(3),s={},p={},d=n.exports=o();d.backlog=s},{}],gos:[function(e,n,t){function r(e,n,t){if(o.call(e,n))return e[n];var r=t();if(Object.defineProperty&&Object.keys)try{return Object.defineProperty(e,n,{value:r,writable:!0,enumerable:!1}),r}catch(i){}return e[n]=r,r}var o=Object.prototype.hasOwnProperty;n.exports=r},{}],handle:[function(e,n,t){function r(e,n,t,r){o.buffer([e],r),o.emit(e,n,t)}var o=e("ee").get("handle");n.exports=r,r.ee=o},{}],id:[function(e,n,t){function r(e){var n=typeof e;return!e||"object"!==n&&"function"!==n?-1:e===window?0:a(e,i,function(){return o++})}var o=1,i="nr@id",a=e("gos");n.exports=r},{}],loader:[function(e,n,t){function r(){if(!E++){var e=x.info=NREUM.info,n=l.getElementsByTagName("script")[0];if(setTimeout(s.abort,3e4),!(e&&e.licenseKey&&e.applicationID&&n))return s.abort();c(y,function(n,t){e[n]||(e[n]=t)}),f("mark",["onload",a()+x.offset],null,"api");var t=l.createElement("script");t.src="https://"+e.agent,n.parentNode.insertBefore(t,n)}}function o(){"complete"===l.readyState&&i()}function i(){f("mark",["domContent",a()+x.offset],null,"api")}function a(){return O.exists&&performance.now?Math.round(performance.now()):(u=Math.max((new Date).getTime(),u))-x.offset}var u=(new Date).getTime(),f=e("handle"),c=e(3),s=e("ee"),p=e(2),d=window,l=d.document,m="addEventListener",v="attachEvent",g=d.XMLHttpRequest,w=g&&g.prototype;NREUM.o={ST:setTimeout,SI:d.setImmediate,CT:clearTimeout,XHR:g,REQ:d.Request,EV:d.Event,PR:d.Promise,MO:d.MutationObserver};var h=""+location,y={beacon:"bam.nr-data.net",errorBeacon:"bam.nr-data.net",agent:"js-agent.newrelic.com/nr-1130.min.js"},b=g&&w&&w[m]&&!/CriOS/.test(navigator.userAgent),x=n.exports={offset:u,now:a,origin:h,features:{},xhrWrappable:b,userAgent:p};e(1),l[m]?(l[m]("DOMContentLoaded",i,!1),d[m]("load",r,!1)):(l[v]("onreadystatechange",o),d[v]("onload",r)),f("mark",["firstbyte",u],null,"api");var E=0,O=e(5)},{}]},{},["loader"]);</script><link rel="stylesheet" href="/cpopg/css/spw/spwConcatenado.css" type="text/css"/>
<link rel="stylesheet" href="/cpopg/css/esaj.css" type="text/css"/>
<link rel="stylesheet" href="/cpopg/css/saj/saj-captcha.css" type="text/css"/>
<link rel="stylesheet" href="/cpopg/css/cpo.css" type="text/css"/>	
<link rel="stylesheet" href="/cpopg/css/formulario.css" type="text/css"/>
<link rel="stylesheet" href="/cpopg/css/saj/saj-popup-modal.css" type="text/css"/>

<link rel="shortcut icon" href="/favicon.ico" type="image/x-icon" />

<script language="javascript" type="text/JavaScript" src="/cpopg/js/jquery/jquery.min.js?n=70d0426f-74bf-47e3-b068-4fff9b2b337b"></script>
<script language="javascript" type="text/JavaScript" src="/cpopg/js/jquery/jquery.func_toggle.js?n=9e2cd0b3-a939-4be1-b9e4-99e8bdb01467"></script>
<script language="javascript" type="text/javascript" src="/cpopg/js/jquery/jquery.jplayer.2.9.2.min.js?n=e277a533-e072-45ff-aa77-e9e57fec6256"></script>
<script language="javascript" type="text/javascript" src="/cpopg/js/jquery/jquery.blockUI.min.js?n=648489c4-5944-4f29-b838-34d5d1cb9a94"></script>
<script language="javascript" type="text/javascript" src="/cpopg/js/jquery/jquery.browser.min.js?n=35765b2d-2619-48c9-bda0-f0988baee504"></script>

<script language="javascript" type="text/JavaScript" src="/cpopg/js/saj/saj-web-2.2.41-4.js?n=77f0902a-3d25-4ef0-85d1-c774b61314bb"></script>
<script language="javascript" type="text/JavaScript" src="/cpopg/js/saj/saj-tooltip.js?n=a24d56ab-7432-4cb8-8aba-ce79d8a7f567"></script>
<script language="javascript" type="text/JavaScript" src="/cpopg/js/saj/saj-popup-modal-1.0.0-1.js?n=7d4068ab-8bf6-4b2d-9644-40efa6403cff"></script>
<script language="javascript" type="text/JavaScript" src="/cpopg/js/saj/saj-browser.js?n=cc5ec8cc-8e5b-4547-8f1c-92bc17fb4cfa"></script>
<script language="javascript" type="text/JavaScript" src="/cpopg/js/saj/saj-mascara-input.js?n=093755b0-890b-439f-823a-e5bf45002ca5"></script>
<script language="javascript" type="text/JavaScript" src="/cpopg/js/saj/saj-numeroProcesso.js?n=9968e07b-7f86-43b1-b7e2-e1754ac576ab"></script>
<script language="javascript" type="text/JavaScript" src="/cpopg/js/spw/spwConcatenado.js?n=eb01d2a0-21a9-4929-9b2a-e0c694fe929d"></script>

 
<script>
	window.saj = window.saj || {};
	window.saj.env = window.saj.env || {};
	window.saj.env.root = '/cpopg';
	window.saj.env.css = '/cpopg/css';
	window.saj.env.js = '/cpopg/js';
	window.saj.env.imagens = '/cpopg/imagens';
	window.saj.env.queryString = 'processo.codigo=01000O7550000&processo.foro=1&uuidCaptcha=sajcaptcha_e7177bac99504aa4b68783240bfba9d8';	

	window.saj.cpo = window.saj.cpo || {};

	// transfere variavel se cpo esta rodando para totem
	window.saj.cpo.totem = 'false' == 'true';
</script>

<script language="javascript" type="text/JavaScript" src="/cpopg/js/saj-cpo-cbpesquisa.js?n=629714429"></script>
<script language="javascript" type="text/JavaScript" src="/cpopg/js/formulario.js?n=1425064359"></script>
<script language="javascript" type="text/JavaScript" src="/cpopg/js/saj/acessibilidade/acessibilidade.js?n=b8e3122c-05e7-4d27-a25d-0672749a76f3"></script>
	




<script>
	(function($){
		$(function(){
			var intervalo = 1795000;
			$.saj.manterSessao('/cpopg/manterSessao.do', intervalo);
		});
				
	})(jQuery);
</script>




	
    <script language="javascript" type="text/JavaScript">
        $.saj = $.saj || {};
        $.saj.acessoRecurso = {
            abrirPastaDigitalEmPopup: 'true' == 'true',
            idRecursoQueProvocouLiberacaoPorSenha: '',
            popupSenha: {
                mostrar: 'false' == 'true',
                titulo: 'Digite a senha do processo',
                altura: 220 + Number("10"),
                alturaAdicionalParaMensagemValidacao: Number("0"),
                largura: 480
            }
        };

        $.saj.getUrlParameter = function getUrlParameter(sParam) {
            var sPageURL = decodeURIComponent(window.location.search.substring(1)), sURLVariables = sPageURL.split('&'), sParameterName, i;
            for (i = 0; i < sURLVariables.length; i++) {
                sParameterName = sURLVariables[i].split('=');
                if (sParameterName[0] === sParam) {
                    return sParameterName[1] === undefined ? true : sParameterName[1];
                }
            }
        };
    </script>
    <script language="javascript" type="text/JavaScript" src="/cpopg/jsp/show-2.8.33-0.js?n=2077829452"></script>

	















	
	
		


<meta http-equiv="Content-Type" content="text/html; charset=iso-8859-1"/>
<meta http-equiv="pragma" content="no-cache"/>
<meta http-equiv="cache-control" content="no-cache"/>
<meta http-equiv="expires" content="0"/>

<link href="https://www2.tjal.jus.br/esaj/tema/padrao/css/esajComum.css" rel="stylesheet" type="text/css" />
<link href="https://www2.tjal.jus.br/esaj/tema/padrao/css/esajLayout.css" rel="stylesheet" type="text/css" />
<link href="https://www2.tjal.jus.br/esaj/tema/padrao/clientes/AL/css/esajLayoutAL.css" rel="stylesheet" type="text/css" />

<style type="text/css">
<!--
/* botao default means o mais claro, que seria o "botï¿½o principal" */
.spwBotaoDefault, .spwBotaoDefault-o { 
	background-image: url(https://www2.tjal.jus.br/esaj/tema/padrao/clientes/AL/imagens/layout/fundoBotaoDefault.gif);
}
/* o botao secundario, menos provavel de ser clicado, mais escuro */
.spwBotao, .spwBotao-o { 
	background-image: url(https://www2.tjal.jus.br/esaj/tema/padrao/imagens/layout/fundoBotao.gif);
}
-->
</style>


<script language="javascript" type="text/JavaScript" src="https://www2.tjal.jus.br/esaj/tema/padrao/js/menu.js?n=2553626853"></script>

<link rel="shortcut icon" href="https://www2.tjal.jus.br/esaj/tema/padrao/clientes/AL/imagens/favicon.ico" type="image/x-icon" />


	


  

	<script type="text/javascript">
		(function($) {
			$(function(){
				// Correção temporária do alinhamento do divisor de seção de formulário do SPW 
				$('td[background$="spw/fundo_subtitulo2.gif"]').attr('valign', 'top');
				
			})
		})(jQuery);
	</script>
</head>


<body>
<div class="div-conteudo">
	
	
	    


































	
	
	        








	




















	
	
		<table width="100%" border="0" cellpadding="0" cellspacing="0" class="esajTabelaCabecalhoCliente" summary=" ">
	<tr>
		<td><a href="http://www.tjal.jus.br"><img src="https://www2.tjal.jus.br/esaj/tema/padrao/clientes/AL/imagens/layout/cabecalhoLogo.jpg" width="433" height="58" alt="Tribunal de Justiça de Alagoas" /></a></td>
		<td align="right" valign="top"><img src="https://www2.tjal.jus.br/esaj/tema/padrao/clientes/AL/imagens/layout/cabecalhoImagem.jpg" alt=" " width="353" height="58" /></td>
	</tr>
</table>
	


  












	
	
		


<!-- cabecalho e-Saj -->
<table width="100%" cellpadding="0" cellspacing="0" class="esajTabelaCabecalhoServico" summary=" ">
	<tr>
		<td class="esajCelulaCabecalhoServicoLogo">




<a href="https://www2.tjal.jus.br/esaj/redirecionarNovoEsaj.do"><img src="https://www2.tjal.jus.br/esaj/tema/padrao/imagens/layout/eSajServico.gif" title="Voltar para página inicial do e-SAJ" alt="Voltar para página inicial do e-SAJ" width="255" height="53" border="0" /></a><img src="https://www2.tjal.jus.br/esaj/tema/padrao/imagens/layout/eSajServicoInf.gif" alt=" " width="255" height="28" border="0" /></td>
		<td class="esajCelulaCabecalhoServicoMenu">
		    <table width="100%" border="0" cellpadding="0" cellspacing="0" summary=" ">
				<tr>
					<td align="right">








<!-- Devido ao alinhamento inline das imagens, elas não podem ter nenhum tipo de espaçamento entre elas, portanto é necessário finalizar a linha iniciando um comentário e fechar o comentário imediatamente antes da abertura da nova tag-->
<img height="24" width="47" alt=" " src="https://www2.tjal.jus.br/esaj/tema/padrao/imagens/layout/inicioMenuSup.jpg"/><a href="https://www2.tjal.jus.br/caixapostal"><!--
--><img height="24" width="83" border="0" alt="Caixa postal" src="https://www2.tjal.jus.br/esaj/tema/padrao/imagens/layout/caixaPostal.gif"/></a><!--
--><img height="24" width="4" alt=" " src="https://www2.tjal.jus.br/esaj/tema/padrao/imagens/layout/menuSeparador.jpg"/><a href="https://www2.tjal.jus.br/esaj/cadastro.do"><!--
--><img src="https://www2.tjal.jus.br/esaj/tema/padrao/imagens/layout/cadastro.gif" alt="Cadastro" width="81" height="24" border="0" /></a><!--


    --><img height="24" width="4" alt=" " src="https://www2.tjal.jus.br/esaj/tema/padrao/imagens/layout/menuSeparador.jpg"/><a target="blank" href="https://www2.tjal.jus.br/WebHelp/"><!--
    --><img height="24" width="61" border="0" alt="Ajuda do e-Saj" src="https://www2.tjal.jus.br/esaj/tema/padrao/imagens/layout/ajuda.gif"/></a><!--


--></td>
				</tr>
				<tr>
					<td class="esajCelulaIdentificacaoServico" >




<style>

    td.esajCelulaIdentificacao, td.esajCelulaIdentificacaoServico {
        position: relative;
    }

    button.escolhaBeta {
        border: 0;
        background: #0078D7;
        color: #fff;
        margin-right: 5px;
        font-family: verdana;
        font-size: 12px;
        height: 27px;
        line-height: 14px;
        border-radius: 0;
        text-shadow: 1px 1px 1px #797979;
        position: absolute;
        top: 30px;
        right: 3px;
        padding: 7px 12px;
        cursor: pointer;
    }

    button.escolhaBeta:hover {
        background-color: #0b5c9c;
    }
</style>

<span id="identificacao"></span>

<button style="display: none" class="escolhaBeta"></button>




    
    
        <script language="javascript" type="text/JavaScript" src="https://www2.tjal.jus.br/sajcas/conteudoIdentificacao?script=1564409788048"></script>
    

</td>
				</tr>
				<tr>
					<td class="esajCelulaCabecalhoServicoCaminho" >









  







  



  







  



  










<a href="" class="esajCaminhoLink"></a>

 &gt;


<a href="https://www2.tjal.jus.br/esaj/portal.do?servico=740000" class="esajCaminhoLink">Bem-vindo</a>

 &gt;


<a href="https://www2.tjal.jus.br/esaj/portal.do?servico=190090" class="esajCaminhoLink">Consultas Processuais</a>

 &gt;



Consulta de Processos de 1º Grau
</td>
				</tr>
			</table>
		</td>
	</tr>
</table>
<!-- CONTEUDO -->
<table width="100%" border="0" cellpadding="0" cellspacing="0" class="esajTabelaTitulo">
	<tr>
		<td class="esajCelulaMenuSuspenso">
		<a href="#" onclick="return showMenuSuspenso()"><img src="https://www2.tjal.jus.br/esaj/tema/padrao/imagens/layout/menuSuspenso.gif" alt="Menu de servi&ccedil;os" width="243" height="21" style="display:block" /></a>
		<div id="layerMenu" style="display:none">
			

<!-- MENU -->
<ul id="esajMenuArea">
  <li class="esajMenuFechado"><a href="#">Consulta de Temas de Repercussão Geral e Casos Repetitivos</a></li>
	<li class="esajMenuFechado"><a href="#" onclick="return esajMenu(this)">Recolhimento de Custas</a>
<ul>
	<li class="esajMenuFechado"><a href="#" onclick="return esajMenu(this)">Custas de 1º Grau</a>
<ul>
  <li class="esajMenuItem"><a href="https://www2.tjal.jus.br/ccpweb/abrirConsultaCustas.do">Consulta de Custas de 1º Grau</a></li>
  <li class="esajMenuItem"><a href="https://www2.tjal.jus.br/ccpweb/iniciarCalculoDeCustas.do?cdTipoCusta=1&flTipoCusta=0&cdServicoCalculoCusta=690003">Custas Iniciais de 1º Grau</a></li>
  <li class="esajMenuItem"><a href="https://www2.tjal.jus.br/ccpweb/iniciarCalculoDeCustas.do?cdTipoCusta=15&flTipoCusta=5&cdServicoCalculoCusta=690004">Atos Avulsos de 1º Grau</a></li>
  <li class="esajMenuItem"><a href="https://www2.tjal.jus.br/ccpweb/iniciarCalculoDeCustas.do?cdTipoCusta=10&flTipoCusta=1&cdServicoCalculoCusta=690005">Custas de Preparo de 1º Grau</a></li>
  <li class="esajMenuItem"><a href="https://www2.tjal.jus.br/ccpweb/iniciarCalculoDeCustas.do?cdTipoCusta=14&flTipoCusta=0&cdServicoCalculoCusta=690006">Custas Juizado Especial - Recurso Inominado</a></li>
  <li class="esajMenuItem"><a href="https://www2.tjal.jus.br/ccpweb/iniciarCalculoDeCustas.do?cdTipoCusta=3&flTipoCusta=1&cdServicoCalculoCusta=690009">Custas Intermediárias de 1º Grau</a></li>
</ul>
  </li>
	<li class="esajMenuFechado"><a href="#" onclick="return esajMenu(this)">Custas Processuais de Segundo Grau</a>
<ul>
  <li class="esajMenuItem"><a href="https://www2.tjal.jus.br/ccpweb/abrirConsultaCustasSg.do">Titulo que será mostrado no portal</a></li>
  <li class="esajMenuItem"><a href="https://www2.tjal.jus.br/ccpweb/iniciarCalculoDeCustasSg.do?cdServicoCalculoCusta=690205&flTipoCusta=0&cdTipoCusta=1">Custas iniciais de 2º Grau</a></li>
</ul>
  </li>
</ul>
  </li>
	<li class="esajMenuAberto"><a href="#" onclick="return esajMenu(this)">Consultas</a>
<ul>
  <li class="esajMenuItem"><a href="https://www2.tjal.jus.br/cpopg/open.do">Processos de 1º Grau</a></li>
	<li class="esajMenuFechado"><a href="#" onclick="return esajMenu(this)">Ordem de Processos</a>
<ul>
  <li class="esajMenuItem"><a href="https://www2.tjal.jus.br/cop/abrirConsultaDeOrdemDeJulgamentoPg.do">Julgamento do 1º Grau</a></li>
  <li class="esajMenuItem"><a href="https://www2.tjal.jus.br/cop/abrirConsultaDeOrdemDeAtoPg.do">Publicação e Cumprimento de Atos de 1º Grau</a></li>
</ul>
  </li>
  <li class="esajMenuItem"><a href="https://www2.tjal.jus.br/cposg5/open.do">Processos de 2º Grau</a></li>
</ul>
  </li>
	<li class="esajMenuFechado"><a href="#" onclick="return esajMenu(this)">Jurisprudência</a>
<ul>
  <li class="esajMenuItem"><a href="https://www2.tjal.jus.br/cjsg/consultaSimples.do">Simples</a></li>
  <li class="esajMenuItem"><a href="https://www2.tjal.jus.br/cjsg/consultaCompleta.do">Completa</a></li>
</ul>
  </li>
  <li class="esajMenuFechado"><a href="https://www2.tjal.jus.br/cdje">Diário da Justiça Eletrônico</a></li>
	<li class="esajMenuFechado"><a href="#" onclick="return esajMenu(this)">Push</a>
<ul>
  <li class="esajMenuItem"><a href="https://www2.tjal.jus.br/pushpg">1º Grau</a></li>
  <li class="esajMenuItem"><a href="https://www2.tjal.jus.br/pushsg3">Push - Segundo Grau - Físico</a></li>
  <li class="esajMenuItem"><a href="https://www2.tjal.jus.br/pushsg5">2º Grau - Digital</a></li>
</ul>
  </li>
	<li class="esajMenuFechado"><a href="#" onclick="return esajMenu(this)">Conferência de Documentos</a>
<ul>
  <li class="esajMenuItem"><a href="https://www2.tjal.jus.br/pastadigital/pg/abrirConferenciaDocumento.do">Documento Digital do 1º Grau</a></li>
  <li class="esajMenuItem"><a href="https://www2.tjal.jus.br/pastadigital/sg/abrirConferenciaDocumento.do">Documento Digital do 2º Grau</a></li>
</ul>
  </li>
	<li class="esajMenuFechado"><a href="#" onclick="return esajMenu(this)">Certidões</a>
<ul>
	<li class="esajMenuFechado"><a href="#" onclick="return esajMenu(this)">Certidões de 1º grau</a>
<ul>
  <li class="esajMenuItem"><a href="https://www2.tjal.jus.br/sco/abrirCadastro.do">Pedido de Certidão de 1º Grau</a></li>
  <li class="esajMenuItem"><a href="https://www2.tjal.jus.br/sco/abrirConferencia.do">Conferência de Certidão de 1º Grau</a></li>
  <li class="esajMenuItem"><a href="https://www2.tjal.jus.br/sco/abrirDownload.do">Baixar Certidão de 1º Grau</a></li>
</ul>
  </li>
</ul>
  </li>
  <li class="esajMenuFechado"><a href="https://www2.tjal.jus.br/cpoesajsg/pcpoOrgaosJulgadores.jsp">Consulta da Pauta de Julgamento - Cível</a></li>
	<li class="esajMenuFechado"><a href="#" onclick="return esajMenu(this)">Intimação e Citação Eletrônica</a>
<ul>
  <li class="esajMenuItem"><a href="https://www2.tjal.jus.br/intimacoesweb/abrirConsultaAtosRecebidos.do">Consulta de Recebidas</a></li>
  <li class="esajMenuItem"><a href="https://www2.tjal.jus.br/intimacoesweb/abrirConsultaAtosNaoRecebidos.do">Recebimento</a></li>
</ul>
  </li>
	<li class="esajMenuFechado"><a href="#" onclick="return esajMenu(this)">Peticionamento Eletrônico</a>
<ul>
	<li class="esajMenuFechado"><a href="#" onclick="return esajMenu(this)">Peticionamento Eletrônico de 1º Grau</a>
<ul>
  <li class="esajMenuItem"><a href="https://www2.tjal.jus.br/petpg/abrirNovaPeticaoInicial.do?instancia=PG">Peticionamento Inicial de 1º Grau</a></li>
  <li class="esajMenuItem"><a href="https://www2.tjal.jus.br/petpg/abrirNovaPeticaoIntermediaria.do?instancia=PG">Peticionamento Intermediário de 1º Grau</a></li>
  <li class="esajMenuItem"><a href="https://www2.tjal.jus.br/petpg/abrirVerificacaoRequisitosPet.do">Verificação de Requisitos</a></li>
  <li class="esajMenuItem"><a href="https://www2.tjal.jus.br/petpg/abrirConsultaPeticoes.do">Consulta de Petições de 1º Grau</a></li>
</ul>
  </li>
	<li class="esajMenuFechado"><a href="#" onclick="return esajMenu(this)">Peticionamento Eletrônico de 2º Grau</a>
<ul>
  <li class="esajMenuItem"><a href="https://www2.tjal.jus.br/petsg/abrirNovaPeticaoInicial.do?instancia=SG">Peticionamento Inicial de 2º Grau</a></li>
  <li class="esajMenuItem"><a href="https://www2.tjal.jus.br/petsg/abrirNovaPeticaoIntermediaria.do?instancia=SG">Peticionamento Intermediário de 2º Grau</a></li>
  <li class="esajMenuItem"><a href="https://www2.tjal.jus.br/petsg/abrirConsultaPeticoes.do">Consulta de Petições de 2º Grau</a></li>
</ul>
  </li>
	<li class="esajMenuFechado"><a href="#" onclick="return esajMenu(this)">Peticionamento Eletrônico de Turmas Recursais / Plantão (2º Grau)/ Precatórios</a>
<ul>
  <li class="esajMenuItem"><a href="https://www2.tjal.jus.br/petcr/abrirNovaPeticaoInicial.do?instancia=CR">Peticionamento Inicial - TR/ Plantão TJ/ Precat.</a></li>
  <li class="esajMenuItem"><a href="https://www2.tjal.jus.br/petcr/abrirNovaPeticaoIntermediaria.do?instancia=CR">Peticionamento Intermediário - TR/ Plantão TJ/ Precat.</a></li>
  <li class="esajMenuItem"><a href="https://www2.tjal.jus.br/petcr/abrirConsultaPeticoes.do">Consulta de Petições - TR/ Plantão TJ/ Precat.</a></li>
</ul>
  </li>
</ul>
  </li>
  <li class="esajMenuFechado"><a href="https://www2.tjal.jus.br/cposg5pauta/pcpoOrgaosJulgadores.jsp">Pauta de Julgamento - Criminal</a></li>
</ul>

		</div>
		</td>
		<td class="esajCelulaTituloServico"><h1 class="esajTituloPagina">Consulta de Processos de 1º Grau</h1></td>
	</tr>
</table>
<table width="100%" cellpadding="0" cellspacing="0" summary=" ">
	<tr>
		<td class="esajCelulaConteudoServico" >
		
	


  



	        	
	        
	            








































    <span class="esajTituloOrientacoes">Orientações</span>

    
        
        
            <ul class="esajUlOrientacoes" id="">
                
    















	
	<li>
		




















		


Processos distribuídos no mesmo dia podem ser localizados se buscados pelo número do processo, com o seu foro selecionado.


		
	</li>



    

















    

















    















	
	<li>
		




















		


Dúvidas? Clique <a style="cursor:pointer" class="layout" onclick="popup('/WebHelp/id_consultas_processuais.htm','','location=no, toolbar=no, resizable=yes, width=795, height=560, scrollbars=yes')">aqui</a> para mais informações sobre como pesquisar.


		
	</li>



    















	
	<li>
		




















		


Processos baixados, em segredo de justiça ou distribuídos no mesmo dia serão apresentados somente na pesquisa pelo número do processo.


		
	</li>




            </ul>
        
    

    
    <br>








    









<form name="consultarProcessoForm" method="GET" action="/cpopg/search.do" autocomplete="off" enctype="" onsubmit="return applySubmit(this, eval('function spwSubmit(t, e){decodificaInputMulSelOnSubmit();if (!IS_enableSubmit) return false; return BENV_isCamposValidos(t); } spwSubmit(this, event);'));" id="formConsulta" target="">
	<input type="hidden" name="conversationId" value="">

        




























	
	
	
	
		
			
		
		
	
	
	<div class="">
		
			<br/>
			
			
				
				
					









<table width="100%" border="0" cellspacing="0" cellpadding="0">
	<tr valign="top">
		<td height="21" valign="top" nowrap background="/cpopg/imagens/spw/fundo_subtitulo.gif"> 
			
				
				
					<h2 class="subtitle">
						Dados para pesquisa
						
					</h2>
				
			
		</td>
		<td background="/cpopg/imagens/spw/fundo_subtitulo2.gif" width="90%" aria-hidden="true">
			<img src="/cpopg/imagens/spw/final_subtitulo.gif" width="16" height="20" tabindex="-1">
		</td>
	</tr>
</table>

				
			
			<br/>
		
		<table id="secaoFormConsulta" class="secaoFormBody" width="100%" style="" cellpadding="2" cellspacing="0" border="0">
			
            
            

            
                
                

































	



























	
	
		
		
	






	
	
		
		
			
				<tr class="">
				
			
		
		
		
		  <td id="" width="150" valign="">
			
			  
		  	  
		      
				
		  	    <label for="id_Comarca" class="" style="text-align:right;font-weight:bold;;">Comarca:</label>
		      
		    
		  </td>
		
		
		
		  
		  
		    <td valign="">
		  
		
			
	
	
	
	
	
	
		
		
		<select name="dadosConsulta.localPesquisa.cdLocal" id="id_Comarca" obrigatorio="" rotulo="Comarca"><option value="-1">



















	
	
					



Todas as comarcas

</option>
                
			<option value="152"> Juizado de São Miguel dos Campos</option>
<option value="73">Corregedoria Geral da Justiça</option>
<option value="10">Foro Colônia Leopoldina</option>
<option value="202">Foro de Água Branca</option>
<option value="203">Foro de Anadia</option>
<option value="58">Foro de Arapiraca</option>
<option value="40">Foro de Atalaia</option>
<option value="204">Foro de Batalha</option>
<option value="5">Foro de Boca da Mata</option>
<option value="6">Foro de Cacimbinhas</option>
<option value="7">Foro de Cajueiro</option>
<option value="8">Foro de Campo Alegre</option>
<option value="9">Foro de Canapi</option>
<option value="41">Foro de Capela</option>
<option value="59">Foro de Chã Preta (Fora de Uso)</option>
<option value="42">Foro de Coruripe</option>
<option value="43">Foro de Delmiro Gouveia</option>
<option value="60">Foro de Feira Grande</option>
<option value="11">Foro de Flexeiras</option>
<option value="12">Foro de Girau do Ponciano</option>
<option value="13">Foro de Igaci</option>
<option value="14">Foro de Igreja Nova</option>
<option value="15">Foro de Joaquim Gomes</option>
<option value="16">Foro de Junqueiro</option>
<option value="17">Foro de Limoeiro de Anadia</option>
<option value="1" selected="selected">Foro de Maceió</option>
<option value="96">Foro de Maceió - 25ª Vara Família</option>
<option value="18">Foro de Major Isidoro</option>
<option value="19">Foro de Maragogi</option>
<option value="20">Foro de Maravilha</option>
<option value="44">Foro de Marechal Deodoro</option>
<option value="21">Foro de Maribondo</option>
<option value="22">Foro de Mata Grande</option>
<option value="23">Foro de Matriz de Camaragibe</option>
<option value="61">Foro de Messias</option>
<option value="45">Foro de Murici</option>
<option value="24">Foro de Novo Lino</option>
<option value="25">Foro de Olho DÁgua das Flores</option>
<option value="46">Foro de Palmeira dos Índios</option>
<option value="48">Foro de Pão de Açúcar</option>
<option value="28">Foro de Paripueira</option>
<option value="27">Foro de Passo de Camaragibe</option>
<option value="29">Foro de Paulo Jacinto</option>
<option value="49">Foro de Penedo</option>
<option value="26">Foro de Piaçabuçu</option>
<option value="47">Foro de Pilar</option>
<option value="30">Foro de Piranhas</option>
<option value="66">Foro de Plantão Cível da Capital</option>
<option value="67">Foro de Plantão Criminal da Capital</option>
<option value="50">Foro de Porto Calvo</option>
<option value="31">Foro de Porto de Pedras</option>
<option value="32">Foro de Porto Real do Colégio</option>
<option value="33">Foro de Quebrangulo</option>
<option value="51">Foro de Rio Largo</option>
<option value="34">Foro de Santa Luzia do Norte</option>
<option value="55">Foro de Santana do Ipanema</option>
<option value="35">Foro de São Brás</option>
<option value="52">Foro de São José da Laje</option>
<option value="36">Foro de São José da Tapera</option>
<option value="54">Foro de São Luiz do Quitunde</option>
<option value="53">Foro de São Miguel dos Campos</option>
<option value="37">Foro de São Sebastião</option>
<option value="64">Foro de Taquarana</option>
<option value="38">Foro de Teotônio Vilela</option>
<option value="39">Foro de Traipu</option>
<option value="56">Foro de União dos Palmares</option>
<option value="57">Foro de Viçosa</option>
<option value="88">Foro Justiça Itinerante</option>
<option value="68">Foro Plantonista da 1ª Circunscrição</option>
<option value="69">Foro Plantonista da 2ª Circunscrição</option>
<option value="70">Foro Plantonista da 3ª Circunscrição</option>
<option value="71">Foro Plantonista da 4ª Circunscrição</option>
<option value="72">Foro Plantonista da 5ª Circunscrição</option>
<option value="148">Juizado  de Santana do Ipanema</option>
<option value="171">Juizado Criminal e do Torcedor de Maceió</option>
<option value="144">Juizado da Fazenda Pública da Capital</option>
<option value="94">Juizado da Violência Doméstica/Familiar</option>
<option value="145">Juizado de Delmiro Gouveia</option>
<option value="146">Juizado de Palmeira dos Índios</option>
<option value="349">Juizado de Penedo</option>
<option value="147">Juizado de Rio Largo</option>
<option value="356">Juizado de União dos Palmares</option>
<option value="91">1 º Juizado Especial Cível da Capital</option>
<option value="149">1º Juizado Cível e Criminal de Arapiraca</option>
<option value="84">1º Vara Infância e Juventude da Capital</option>
<option value="81">10º Juizado Especial Cível da Capital</option>
<option value="80">11º Juizado Especial Cível da Capital</option>
<option value="143">12º Juizado Cível e Criminal da Capital</option>
<option value="150">2º Juizado Cível e Criminal de Arapiraca</option>
<option value="92">2º Juizado Especial Cível da Capital</option>
<option value="93">26º Vara Civel da Capital/Familia (Fora</option>
<option value="90">28ª Vara Infância e Juventude da Capital</option>
<option value="95">29º Vara Cível da Capital-Conflitos Agrá</option>
<option value="78">3º Juizado Especial Cível da Capital</option>
<option value="205">5º Juizado Especial Cível da Capital</option>
<option value="97">5ª Vara Criminal da Capital - Regional</option>
<option value="75">6º Juizado Especial Cível da Capital</option>
<option value="76">7º Juizado Especial Cível da Capital</option>
<option value="77">8º Juizado Especial Cível da Capital</option>
<option value="82">9º Juizado Especial Cível da Capital</option></select>
	
	
	
	
	
	
		
		</td>
		
			</tr>
		
	
	



            

            
            











































	
	
		
		
	






	
	
		
		
			
				<tr class="">
				
			
		
		
		
		  <td id="" width="150" valign="">
			
			  
		  	  
		      
				
		  	    <label for="cbPesquisa" class="" style="text-align:right;font-weight:bold;;">Pesquisar por:</label>
		      
		    
		  </td>
		
		
		
		  
		  
		    <td valign="">
		  
		
			
	
	
	
	
	
	<table cellpadding="0" cellspacing="0" border="0" width="">
		<tr>
			<td>
				
				
				
                









<select name="cbPesquisa" id="cbPesquisa"><option value="NUMPROC" selected="selected">Número do Processo</option>	
	
 		<option value="NMPARTE">Nome da parte</option>	
	
 		<option value="DOCPARTE">Documento da Parte</option>	
	
 		<option value="NMADVOGADO">Nome do Advogado</option>	
	
 		<option value="NUMOAB">OAB</option>	
	
 		<option value="PRECATORIA">Nº da Carta Precatória na Origem</option>	
	
 		<option value="DOCDELEG">Nº do Documento na Delegacia</option></select>

            
				
			</td>
		</tr>
	  </table>
	
	
	
	
	
	
		
		</td>
		
			</tr>
		
	
	




            
            
                
                    
                    
                        
                            
                            
                                








 
























	
	
	
	
	
	
	

	
	 
		
	
	 
		
	
	
	
	
	
	
		<script>$.saj.numeroProcesso.desabilitarNumeroOculto = false;</script>	
	
	
		
	

	
		
		
			



















<script type="text/javascript">
	(function($){
		$(function() {
			if($.browser.msie){
				$('.grupoRadio').find('input:radio:visible:enabled').attr('aria-required','');
				return;
			}
			$('.grupoRadio').attr('aria-required','').attr('required','');
		});
	})(jQuery);
</script>


























	
	
		
		
	






	
	
		
		
			
				<tr class="">
				
			
		
		
		
		  <td id="" width="150" valign="">
			
		  </td>
		
		
		
		  
		  
		    <td valign="">
		  
		
			
	
	
	
	
	
	<table cellspacing="0" border="0" width="">
		<tr>
			<td>
				
				<fieldset class="grupoRadio" style="margin:0;padding:0;border:none" >
					<legend style="display:block;font-size:0;">
						Tipo do número
					</legend>
					
			   <input type="radio" name="dadosConsulta.tipoNuProcesso" value="UNIFICADO" checked="checked" id="radioNumeroUnificado"><label for="radioNumeroUnificado">Unificado</label>
			   <input type="radio" name="dadosConsulta.tipoNuProcesso" value="SAJ" id="radioNumeroAntigo"><label for="radioNumeroAntigo">Outros</label>
			
				</fieldset>				
				
			</td>
		</tr>
	  </table>
	
	
	
	
	
	
		
		</td>
		
			</tr>
		
	
	



			











































	
	
		
		
	






	
	
		
		
			
				
				<tr id="NUMPROC" class="">
			
		
		
		
		  <td id="" width="150" valign="">
			
			  
		  	  
		      
				
		  	    <label for="" class="" style="text-align:right;font-weight:bold;;">Número do Processo:</label>
		      
		    
		  </td>
		
		
		
		  
		  
		    <td valign="">
		  
		
			
	
	
	
	
	
	<table cellpadding="0" cellspacing="0" border="0" width="">
		<tr>
			<td>
				
				
				
				




















	
	
	
	
	
	
	





























	
	






	
	
	
	
	
		<span id="linhaProcessoUnificado">
		
	
	
	








	


	
	
	
	
	
	





 
<input type="text" name="numeroDigitoAnoUnificado" maxlength="25" size="20" value="0710802-55.2018" formatType="TEXT" onkeypress="CT_KPS(this, event);" onblur="CT_BLR(this);" onkeydown="CT_KDN(this, event);" onmousemove="CT_MMOV(this, event);" onmouseout="CT_MOUT(this, event);" onmouseover="CT_MOV(this, event);" onfocus="C_OFC(this, event);" id="numeroDigitoAnoUnificado">
<input type="text" id="JTRNumeroUnificado" size="3" disabled="disabled" value="8.02"/>
<input type="text" name="foroNumeroUnificado" maxlength="4" size="3" value="0001" formatType="TEXT" onkeypress="CT_KPS(this, event);" onblur="CT_BLR(this);" onkeydown="CT_KDN(this, event);" onmousemove="CT_MMOV(this, event);" onmouseout="CT_MOUT(this, event);" onmouseover="CT_MOV(this, event);" onfocus="C_OFC(this, event);" id="foroNumeroUnificado">
<input type="hidden" name="dadosConsulta.valorConsultaNuUnificado" value="0710802-55.2018.8.02.0001" id="nuProcessoUnificadoFormatado">


<script>
	(function($){
		$(function() {
			var saj = $.saj;
			var mascaraNumeroUnificado = saj.mascaraNumeroUnificado;
			var $campoNumeroDigitoAnoUnificado = $('#numeroDigitoAnoUnificado');
			var $campoForoNumeroUnificado = $('#foroNumeroUnificado');
			var $campoJTRNumeroUnificado = $('#JTRNumeroUnificado');
			saj.configurarMascaraInput($campoForoNumeroUnificado, '0000');
			saj.configurarMascaraInput($campoNumeroDigitoAnoUnificado, mascaraNumeroUnificado, $campoForoNumeroUnificado,{
				afterPaste: afterPaste			
			});
			saj.bindCamposNumeroProcesso('');
			
				var conteudoToolTip = '';
				if(!conteudoToolTip){
					conteudoToolTip = '<div style="width: 500px;" ><div><b>Número de Processo Unificado</b></div><div>O sistema disponibiliza facilidades no preenchimento do número unificado, seu formato é NNNNNNN-DD.AAAA.J.TR.OOOO:</div><div style="padding-top: 3px;"><b>NNNNNNN</b>: Caso o número possua zeros à esquerda o sistema preenche-os automaticamente, basta informar o número e o dígito "-" ou ".". Exemplo: ao informar "310-" o sistema irá preencher "0000310-".</div><div><b>DD</b>     : Deve ser preenchido pelo usuário.</div><div><b>AAAA</b>   : Ao informar dois dígitos para o ano o sistema completa o mesmo, basta pressionar a tecla Tab. Exemplo: ao informar "08" e "Tab" o sistema irá preencher "2008".</div><div><b>J.TR</b>    : São números fixos preenchidos pelo sistema.Exemplo: 8.99.</div><div><b>OOOO</b>   : Caso o número possua zeros à esquerda o sistema preenche-os automaticamente, basta informar o número pressionar a tecla Tab. Exemplo: ao informar "10" e "Tab" o sistema irá preencher "0010".</div></div>';
				}
				$('#numeroDigitoAnoUnificado, #foroNumeroUnificado').registrarTooltip({
					conteudoTooltip: conteudoToolTip,
					localImagensTooltip: '/cpopg/imagens/saj',
					offsetHorizontalExtra: '',
					offsetVertficalExtra: '',
					urlConteudoTooltip: '',
					posicaoTooltip: '',
					objReferenciaPosicaoTooltip:$campoForoNumeroUnificado
				});
			
		});		
		
		var afterPaste = function(textoOriginal) {
			var texto = getDigits(textoOriginal);
			if (!texto || texto.length < 16) {
				return;
			}
			
			var jtr = $('#JTRNumeroUnificado').val();
			jtr = getDigits(jtr);
			var jtrDigitado = texto.substring(13,16);
			var foroUnificado;
			//pega os 4 ultimos digitos, e caso o jtr colado for o mesmo, coloca os 4 no campo de foro unificado
			//se o jtr for diferente nao coloca nada no campo de foro unificado
			if (jtr == jtrDigitado) {
				foroUnificado = texto.substring(16,texto.length);
			} else {
				foroUnificado = '';
			}
			
			var numeroDigitoAno = texto.substring(0,13);
			var $numeroDigitoAno = $('#numeroDigitoAnoUnificado');
			$numeroDigitoAno.val(numeroDigitoAno);
			
			var $foroNumeroUnificado = $('#foroNumeroUnificado');
			$foroNumeroUnificado.val(foroUnificado);
			$foroNumeroUnificado.trigger('blur');
			$foroNumeroUnificado.focus();
		};
		
		var getDigits = function(texto) {
			return texto.replace(/[^0-9]/g, '');
		};
		$('#numeroDigitoAnoUnificado').attr('aria-label','Número do processo. Informe os treze primeiros dígitos do número do processo.');
		$('#foroNumeroUnificado').attr('aria-label','Informe os quatro últimos dígitos do número do processo.');
	})(jQuery);
	
</script>

	
	
	
	    
    
	
	
		
	
	
	
	
		</span>
	



				
					
					
						






















 












































	



	


<style>
	/* remove borda vermelha de elementos required */
	input:invalid {
		box-shadow:none;
	}
</style>

<script type="text/javascript">
	(function($){
		$(function() {
			var saj = $.saj;
			var id = 'nuProcessoAntigoFormatado';
			var idObjetoReferencia = '#' + id;
			if ('' !== '' && !false) {
				$(idObjetoReferencia).registrarTooltip({
					conteudoTooltip: '',
					posicaoTooltip: 'direita',
					objReferenciaPosicaoTooltip:$(idObjetoReferencia)
				});
			}
			//remove comportamento de exibir mensagem de erro típica do html5
			$('form:visible').attr('novalidate','');
			if (''){
				$(idObjetoReferencia).attr('aria-required','true').attr('required','');
			}
		});		
	})(jQuery);
	
</script>


























	
	






	
	
	
	
	
		<span id="linhaProcessoAntigo">
		
	
	
	<input type="text" name="dadosConsulta.valorConsulta" size="" value="" formatType="TEXT" formato="25" obrigatorio="" rotulo="" onkeypress="CT_KPS(this, event);" onblur="CT_BLR(this);" onkeydown="CT_KDN(this, event);" onmousemove="CT_MMOV(this, event);" onmouseout="CT_MOUT(this, event);" onmouseover="CT_MOV(this, event);" onfocus="C_OFC(this, event);" disabled="disabled" style="" class="spwCampoTexto disabled " id="nuProcessoAntigoFormatado" title="" alt="">
		
	
	
	
		
	    
    
	
	
	
		
	
	
	
	
		</span>
	





					
				
				
			
				
			</td>
		</tr>
	  </table>
	
	
	
	
	
	
		
		</td>
		
			</tr>
		
	
	



		
	

                            
                        
                    
                    
                    
                    
                
            
                
                    
                    
                    
                    
                        











































	
	
		
		
	






	
	
		
		
			
				
				<tr id="NMPARTE" class="">
			
		
		
		
		  <td id="" width="150" valign="">
			
			  
		  	  
		      
				
		  	    <label for="chNmCompleto" class="" style="text-align:right;font-weight:bold;;">Nome da parte:</label>
		      
		    
		  </td>
		
		
		
		  
		  
		    <td valign="">
		  
		
			
	
	
	
	
	
	<table cellpadding="0" cellspacing="0" border="0" width="">
		<tr>
			<td>
				
				
				
							<label for="campo_NMPARTE" style="display:block;font-size:0;">Nome da parte</label>
							






















 





	








































	



	


<style>
	/* remove borda vermelha de elementos required */
	input:invalid {
		box-shadow:none;
	}
</style>

<script type="text/javascript">
	(function($){
		$(function() {
			var saj = $.saj;
			var id = 'campo_NMPARTE';
			var idObjetoReferencia = '#' + id;
			if ('' !== '' && !false) {
				$(idObjetoReferencia).registrarTooltip({
					conteudoTooltip: '',
					posicaoTooltip: 'direita',
					objReferenciaPosicaoTooltip:$(idObjetoReferencia)
				});
			}
			//remove comportamento de exibir mensagem de erro típica do html5
			$('form:visible').attr('novalidate','');
			if (''){
				$(idObjetoReferencia).attr('aria-required','true').attr('required','');
			}
		});		
	})(jQuery);
	
</script>


























	
	






	
	
	
	
	
		<span id="">
		
	
	
	<input type="text" name="dadosConsulta.valorConsulta" size="" value="" formatType="TEXT" formato="100" obrigatorio="" rotulo="" onkeypress="CT_KPS(this, event);" onblur="CT_BLR(this);" onkeydown="CT_KDN(this, event);" onmousemove="CT_MMOV(this, event);" onmouseout="CT_MOUT(this, event);" onmouseover="CT_MOV(this, event);" onfocus="C_OFC(this, event);" disabled="disabled" style="" class="spwCampoTexto disabled spwCampoTexto" id="campo_NMPARTE" title="" alt="">
		
	
	
	
		
	    
    
	
	
	
		
	
	
	
	
		</span>
	





							















	



<span id="">
	<label for="id_tag.checkbox.nomecompleto.rotulo">
		<input type="checkbox" name="chNmCompleto" value="true" style="vertical-align:middle;" id="id_tag.checkbox.nomecompleto.rotulo">
		Pesquisar por nome completo
	</label>
</span>


						
				
			</td>
		</tr>
	  </table>
	
	
	
	
	
	
		
		</td>
		
			</tr>
		
	
	



                    
                    
                
            
                
                    
                    
                    
                    
                    
                       











































	
	
		
		
	






	
	
		
		
			
				
				<tr id="DOCPARTE" class="">
			
		
		
		
		  <td id="" width="150" valign="">
			
			  
		  	  
		      
				
		  	    <label for="chNmCompleto" class="" style="text-align:right;font-weight:bold;;">Documento da Parte:</label>
		      
		    
		  </td>
		
		
		
		  
		  
		    <td valign="">
		  
		
			
	
	
	
	
	
	<table cellpadding="0" cellspacing="0" border="0" width="">
		<tr>
			<td>
				
				
				
							<label for="campo_DOCPARTE" style="display:block;font-size:0;">Documento da Parte</label>
							






















 





	








































	



	


<style>
	/* remove borda vermelha de elementos required */
	input:invalid {
		box-shadow:none;
	}
</style>

<script type="text/javascript">
	(function($){
		$(function() {
			var saj = $.saj;
			var id = 'campo_DOCPARTE';
			var idObjetoReferencia = '#' + id;
			if ('' !== '' && !false) {
				$(idObjetoReferencia).registrarTooltip({
					conteudoTooltip: '',
					posicaoTooltip: 'direita',
					objReferenciaPosicaoTooltip:$(idObjetoReferencia)
				});
			}
			//remove comportamento de exibir mensagem de erro típica do html5
			$('form:visible').attr('novalidate','');
			if (''){
				$(idObjetoReferencia).attr('aria-required','true').attr('required','');
			}
		});		
	})(jQuery);
	
</script>


























	
	






	
	
	
	
	
		<span id="">
		
	
	
	<input type="text" name="dadosConsulta.valorConsulta" size="" value="" formatType="TEXT" formato="100" obrigatorio="" rotulo="" onkeypress="CT_KPS(this, event);" onblur="CT_BLR(this);" onkeydown="CT_KDN(this, event);" onmousemove="CT_MMOV(this, event);" onmouseout="CT_MOUT(this, event);" onmouseover="CT_MOV(this, event);" onfocus="C_OFC(this, event);" disabled="disabled" style="" class="spwCampoTexto disabled " id="campo_DOCPARTE" title="" alt="">
		
	
	
	
		
	    
    
	
	
	
		
	
	
	
	
		</span>
	





						
				
			</td>
		</tr>
	  </table>
	
	
	
	
	
	
		
		</td>
		
			</tr>
		
	
	



                    
                
            
                
                    
                    
                    
                    
                        











































	
	
		
		
	






	
	
		
		
			
				
				<tr id="NMADVOGADO" class="">
			
		
		
		
		  <td id="" width="150" valign="">
			
			  
		  	  
		      
				
		  	    <label for="chNmCompleto" class="" style="text-align:right;font-weight:bold;;">Nome do Advogado:</label>
		      
		    
		  </td>
		
		
		
		  
		  
		    <td valign="">
		  
		
			
	
	
	
	
	
	<table cellpadding="0" cellspacing="0" border="0" width="">
		<tr>
			<td>
				
				
				
							<label for="campo_NMADVOGADO" style="display:block;font-size:0;">Nome do Advogado</label>
							






















 





	








































	



	


<style>
	/* remove borda vermelha de elementos required */
	input:invalid {
		box-shadow:none;
	}
</style>

<script type="text/javascript">
	(function($){
		$(function() {
			var saj = $.saj;
			var id = 'campo_NMADVOGADO';
			var idObjetoReferencia = '#' + id;
			if ('' !== '' && !false) {
				$(idObjetoReferencia).registrarTooltip({
					conteudoTooltip: '',
					posicaoTooltip: 'direita',
					objReferenciaPosicaoTooltip:$(idObjetoReferencia)
				});
			}
			//remove comportamento de exibir mensagem de erro típica do html5
			$('form:visible').attr('novalidate','');
			if (''){
				$(idObjetoReferencia).attr('aria-required','true').attr('required','');
			}
		});		
	})(jQuery);
	
</script>


























	
	






	
	
	
	
	
		<span id="">
		
	
	
	<input type="text" name="dadosConsulta.valorConsulta" size="" value="" formatType="TEXT" formato="100" obrigatorio="" rotulo="" onkeypress="CT_KPS(this, event);" onblur="CT_BLR(this);" onkeydown="CT_KDN(this, event);" onmousemove="CT_MMOV(this, event);" onmouseout="CT_MOUT(this, event);" onmouseover="CT_MOV(this, event);" onfocus="C_OFC(this, event);" disabled="disabled" style="" class="spwCampoTexto disabled spwCampoTexto" id="campo_NMADVOGADO" title="" alt="">
		
	
	
	
		
	    
    
	
	
	
		
	
	
	
	
		</span>
	





							















	



<span id="">
	<label for="id_tag.checkbox.nomecompleto.rotulo">
		<input type="checkbox" name="chNmCompleto" value="true" style="vertical-align:middle;" id="id_tag.checkbox.nomecompleto.rotulo">
		Pesquisar por nome completo
	</label>
</span>


						
				
			</td>
		</tr>
	  </table>
	
	
	
	
	
	
		
		</td>
		
			</tr>
		
	
	



                    
                    
                
            
                
                    
                    
                    
                    
                    
                       











































	
	
		
		
	






	
	
		
		
			
				
				<tr id="NUMOAB" class="">
			
		
		
		
		  <td id="" width="150" valign="">
			
			  
		  	  
		      
				
		  	    <label for="chNmCompleto" class="" style="text-align:right;font-weight:bold;;">OAB:</label>
		      
		    
		  </td>
		
		
		
		  
		  
		    <td valign="">
		  
		
			
	
	
	
	
	
	<table cellpadding="0" cellspacing="0" border="0" width="">
		<tr>
			<td>
				
				
				
							<label for="campo_NUMOAB" style="display:block;font-size:0;">OAB</label>
							






















 





	








































	



	


<style>
	/* remove borda vermelha de elementos required */
	input:invalid {
		box-shadow:none;
	}
</style>

<script type="text/javascript">
	(function($){
		$(function() {
			var saj = $.saj;
			var id = 'campo_NUMOAB';
			var idObjetoReferencia = '#' + id;
			if ('' !== '' && !false) {
				$(idObjetoReferencia).registrarTooltip({
					conteudoTooltip: '',
					posicaoTooltip: 'direita',
					objReferenciaPosicaoTooltip:$(idObjetoReferencia)
				});
			}
			//remove comportamento de exibir mensagem de erro típica do html5
			$('form:visible').attr('novalidate','');
			if (''){
				$(idObjetoReferencia).attr('aria-required','true').attr('required','');
			}
		});		
	})(jQuery);
	
</script>


























	
	






	
	
	
	
	
		<span id="">
		
	
	
	<input type="text" name="dadosConsulta.valorConsulta" size="" value="" formatType="TEXT" formato="100" obrigatorio="" rotulo="" onkeypress="CT_KPS(this, event);" onblur="CT_BLR(this);" onkeydown="CT_KDN(this, event);" onmousemove="CT_MMOV(this, event);" onmouseout="CT_MOUT(this, event);" onmouseover="CT_MOV(this, event);" onfocus="C_OFC(this, event);" disabled="disabled" style="" class="spwCampoTexto disabled " id="campo_NUMOAB" title="" alt="">
		
	
	
	
		
	    
    
	
	
	
		
	
	
	
	
		</span>
	





						
				
			</td>
		</tr>
	  </table>
	
	
	
	
	
	
		
		</td>
		
			</tr>
		
	
	



                    
                
            
                
                    
                    
                    
                    
                    
                       











































	
	
		
		
	






	
	
		
		
			
				
				<tr id="PRECATORIA" class="">
			
		
		
		
		  <td id="" width="150" valign="">
			
			  
		  	  
		      
				
		  	    <label for="chNmCompleto" class="" style="text-align:right;font-weight:bold;;">Nº da precatória:</label>
		      
		    
		  </td>
		
		
		
		  
		  
		    <td valign="">
		  
		
			
	
	
	
	
	
	<table cellpadding="0" cellspacing="0" border="0" width="">
		<tr>
			<td>
				
				
				
							<label for="campo_PRECATORIA" style="display:block;font-size:0;">Nº da precatória</label>
							






















 





	








































	



	


<style>
	/* remove borda vermelha de elementos required */
	input:invalid {
		box-shadow:none;
	}
</style>

<script type="text/javascript">
	(function($){
		$(function() {
			var saj = $.saj;
			var id = 'campo_PRECATORIA';
			var idObjetoReferencia = '#' + id;
			if ('' !== '' && !false) {
				$(idObjetoReferencia).registrarTooltip({
					conteudoTooltip: '',
					posicaoTooltip: 'direita',
					objReferenciaPosicaoTooltip:$(idObjetoReferencia)
				});
			}
			//remove comportamento de exibir mensagem de erro típica do html5
			$('form:visible').attr('novalidate','');
			if (''){
				$(idObjetoReferencia).attr('aria-required','true').attr('required','');
			}
		});		
	})(jQuery);
	
</script>


























	
	






	
	
	
	
	
		<span id="">
		
	
	
	<input type="text" name="dadosConsulta.valorConsulta" size="" value="" formatType="TEXT" formato="100" obrigatorio="" rotulo="" onkeypress="CT_KPS(this, event);" onblur="CT_BLR(this);" onkeydown="CT_KDN(this, event);" onmousemove="CT_MMOV(this, event);" onmouseout="CT_MOUT(this, event);" onmouseover="CT_MOV(this, event);" onfocus="C_OFC(this, event);" disabled="disabled" style="" class="spwCampoTexto disabled " id="campo_PRECATORIA" title="" alt="">
		
	
	
	
		
	    
    
	
	
	
		
	
	
	
	
		</span>
	





						
				
			</td>
		</tr>
	  </table>
	
	
	
	
	
	
		
		</td>
		
			</tr>
		
	
	



                    
                
            
                
                    
                    
                    
                    
                    
                       











































	
	
		
		
	






	
	
		
		
			
				
				<tr id="DOCDELEG" class="">
			
		
		
		
		  <td id="" width="150" valign="">
			
			  
		  	  
		      
				
		  	    <label for="chNmCompleto" class="" style="text-align:right;font-weight:bold;;">Nº na delegacia:</label>
		      
		    
		  </td>
		
		
		
		  
		  
		    <td valign="">
		  
		
			
	
	
	
	
	
	<table cellpadding="0" cellspacing="0" border="0" width="">
		<tr>
			<td>
				
				
				
							<label for="campo_DOCDELEG" style="display:block;font-size:0;">Nº na delegacia</label>
							






















 





	








































	



	


<style>
	/* remove borda vermelha de elementos required */
	input:invalid {
		box-shadow:none;
	}
</style>

<script type="text/javascript">
	(function($){
		$(function() {
			var saj = $.saj;
			var id = 'campo_DOCDELEG';
			var idObjetoReferencia = '#' + id;
			if ('' !== '' && !false) {
				$(idObjetoReferencia).registrarTooltip({
					conteudoTooltip: '',
					posicaoTooltip: 'direita',
					objReferenciaPosicaoTooltip:$(idObjetoReferencia)
				});
			}
			//remove comportamento de exibir mensagem de erro típica do html5
			$('form:visible').attr('novalidate','');
			if (''){
				$(idObjetoReferencia).attr('aria-required','true').attr('required','');
			}
		});		
	})(jQuery);
	
</script>


























	
	






	
	
	
	
	
		<span id="">
		
	
	
	<input type="text" name="dadosConsulta.valorConsulta" size="" value="" formatType="TEXT" formato="100" obrigatorio="" rotulo="" onkeypress="CT_KPS(this, event);" onblur="CT_BLR(this);" onkeydown="CT_KDN(this, event);" onmousemove="CT_MMOV(this, event);" onmouseout="CT_MOUT(this, event);" onmouseover="CT_MOV(this, event);" onfocus="C_OFC(this, event);" disabled="disabled" style="" class="spwCampoTexto disabled " id="campo_DOCDELEG" title="" alt="">
		
	
	
	
		
	    
    
	
	
	
		
	
	
	
	
		</span>
	





						
				
			</td>
		</tr>
	  </table>
	
	
	
	
	
	
		
		</td>
		
			</tr>
		
	
	



                    
                
            

            
            














<input id="uuidCaptcha" type="hidden" name="uuidCaptcha" value="sajcaptcha_e7177bac99504aa4b68783240bfba9d8"/>




            
            

        
		</table>
	</div>
	
	
	


        











	<table id="" class="secaoBotoesBody" width="100%" style="" cellpadding="2" cellspacing="0" border="0">
		<tr>
 			
 				<td width="150">&nbsp;</td>
 			
			<td align="">
			
            













	
		
	
	<input type="submit" name="pbEnviar" value="Pesquisar" onclick="" onmouseover="B_mOver(this);" onmouseout="B_mOut(this);" onmouseover="B_mOver(this);" onmouseout="B_mOut(this);" class="spwBotaoDefault " id="pbEnviar">


        
			</td>
		</tr>
	</table>


    
</form>




	        
	
	        
	            
	                
	                    



































<script type="text/javascript">
    (function ($) {
        $(function () {
            var captcha = $.saj.getUrlParameter('uuidCaptcha');

            if(!captcha){
                return;
            }
            var $processoPrinc = $('.processoPrinc');
            $processoPrinc.attr('href', $processoPrinc.attr('href') + '&uuidCaptcha=' + captcha);

            var $processoPaiApenso = $('.processoPaiApenso');
            $processoPaiApenso.attr('href', $processoPaiApenso.attr('href') + '&uuidCaptcha=' + captcha);

        })
    })(jQuery);
</script>

<!-- pasta digital -->

    <table width="100%" border="0" cellspacing="0" cellpadding="1">
        <tr>
            <td align="left" valign="middle" style="padding-left: 5px;">

                
                    
                    
                        
                    
                

                
                    
                    
                        
                        <a class="linkPasta" id="linkPasta" title="Pasta digital" href="#liberarAutoPorSenha">
                            <img src="/cpopg/imagens/icoPeticionamento.gif" border="0" style="vertical-align: -60%"/>
                            Este processo é digital. Clique aqui para visualizar os autos.
                        </a>
                    
                

            </td>
        </tr>
    </table>

































	
	
	
	
		
			
		
		
	
	
	<div class="">
		
			<br/>
			
			
				
				
					









<table width="100%" border="0" cellspacing="0" cellpadding="0">
	<tr valign="top">
		<td height="21" valign="top" nowrap background="/cpopg/imagens/spw/fundo_subtitulo.gif"> 
			
				
				
					<h2 class="subtitle">
						Dados do processo
						
					</h2>
				
			
		</td>
		<td background="/cpopg/imagens/spw/fundo_subtitulo2.gif" width="90%" aria-hidden="true">
			<img src="/cpopg/imagens/spw/final_subtitulo.gif" width="16" height="20" tabindex="-1">
		</td>
	</tr>
</table>

				
			
			<br/>
		
		<table id="" class="secaoFormBody" width="100%" style="" cellpadding="2" cellspacing="0" border="0">
			

    
    
        











































	
	
		
		
	






	
	
		
		
			
				<tr class="">
				
			
		
		
		
		  <td id="" width="150" valign="">
			
			  
		  	  
		      
				
		  	    <label for="dadosFmt.numero" class="labelClass" style="text-align:right;font-weight:bold;;">Processo:</label>
		      
		    
		  </td>
		
		
		
		  
		  
		    <td valign="">
		  
		
			
	
	
	
	
	
	<table cellpadding="0" cellspacing="0" border="0" width="">
		<tr>
			<td>
				
				
				
            



<!-- Atributos -->









	
	
		<span class="">
			0710802-55.2018.8.02.0001
		</span>
		<span class="" >
			
		</span>
	


            
            
            
        
				
			</td>
		</tr>
	  </table>
	
	
	
	
	
	
		
		</td>
		
			</tr>
		
	
	



    
    
    











































	
	








    
    











































	
	
		
		
	






	
	
		
		
			
				<tr class="">
				
			
		
		
		
		  <td id="" width="150" valign="">
			
			  
		  	  
		      
				
		  	    <label for="numeroProcessoSeDependente" class="labelClass" style="text-align:right;font-weight:bold;;">Classe:</label>
		      
		    
		  </td>
		
		
		
		  
		  
		    <td valign="">
		  
		
			
	
	
	
	
	
	<table cellpadding="0" cellspacing="0" border="0" width="">
		<tr>
			<td>
				
				
				
        


































































	
	






	
	
	
	
	
		<span id="">
		
	
	
	<span id="" class="">Procedimento Ordinário</span>
	
	
	
	    
	    
    
    
	
	
		
	
	
	
	
		</span>
	



&nbsp;


































































	
	






	
	
	
	
	
		<span id="">
		
	
	
	<span id="" class="">&nbsp;</span>
	
	
	
	    
	    
    
    
	
	
		
	
	
	
	
		</span>
	




        
        
    
				
			</td>
		</tr>
	  </table>
	
	
	
	
	
	
		
		</td>
		
			</tr>
		
	
	



    











































	
	






	
	
		
		
			
				<tr class="">
				
			
		
		
		
		  <td id="" width="150" valign="">
			
		  </td>
		
		
		
		  
		  
		    <td valign="">
		  
		
			
	
	
	
	
	
	<table cellpadding="0" cellspacing="0" border="0" width="">
		<tr>
			<td>
				
				
				
        <span class="labelClass">Área:</span> Cível 
    
				
			</td>
		</tr>
	  </table>
	
	
	
	
	
	
		
		</td>
		
			</tr>
		
	
	




    
    


































































	
	
		
		
	






	
	
		
		
			
				<tr class="">
				
			
		
		
		
		  <td id="" width="150" valign="">
			
			  
		  	  
		      
				
		  	    <label for="" class="labelClass" style="text-align:right;font-weight:bold;;">Assunto:</label>
		      
		    
		  </td>
		
		
		
		  
		  
		    <td valign="">
		  
		
			
	
	
	
	
	
	<span id="" class="">Dano Material</span>
	
	
	
	    
	    
    
    
	
	
	
		
		</td>
		
			</tr>
		
	
	





    
    


































































	
	
		
		
	










    
    


































































	
	
		
		
	






	
	
		
		
			
				<tr class="">
				
			
		
		
		
		  <td id="" width="150" valign="">
			
			  
		  	  
		      
				
		  	    <label for="" class="labelClass" style="text-align:right;font-weight:bold;;">Outros assuntos:</label>
		      
		    
		  </td>
		
		
		
		  
		  
		    <td valign="">
		  
		
			
	
	
	
	
	
	<span id="" class="">Dano Moral</span>
	
	
	
	    
	    
    
    
	
	
	
		
		</td>
		
			</tr>
		
	
	






    
    


































































	
	
		
		
	






	
	
		
		
			
				<tr class="">
				
			
		
		
		
		  <td id="" width="150" valign="">
			
			  
		  	  
		      
				
		  	    <label for="" class="labelClass" style="text-align:right;font-weight:bold;;">Distribuição:</label>
		      
		    
		  </td>
		
		
		
		  
		  
		    <td valign="">
		  
		
			
	
	
	
	
	
	<span id="" class="">02/05/2018 às 19:01 - Sorteio</span>
	
	
	
	    
	    
    
    
	
	
	
		
		</td>
		
			</tr>
		
	
	




    


































































	
	






	
	
		
		
			
				<tr class="">
				
			
		
		
		
		  <td id="" width="150" valign="">
			
		  </td>
		
		
		
		  
		  
		    <td valign="">
		  
		
			
	
	
	
	
	
	<span id="" class="">4ª Vara Cível da Capital - Foro de Maceió</span>
	
	
	
	    
	    
    
    
	
	
	
		
		</td>
		
			</tr>
		
	
	





    
    











































	
	
		
		
	









    
    


































































	
	
		
		
	






	
	
		
		
			
				<tr class="">
				
			
		
		
		
		  <td id="" width="150" valign="">
			
			  
		  	  
		      
				
		  	    <label for="" class="labelClass" style="text-align:right;font-weight:bold;;">Controle:</label>
		      
		    
		  </td>
		
		
		
		  
		  
		    <td valign="">
		  
		
			
	
	
	
	
	
	<span id="" class="">2018/000520</span>
	
	
	
	    
	    
    
    
	
	
	
		
		</td>
		
			</tr>
		
	
	





    
    


































































	
	
		
		
	






	
	
		
		
			
				<tr class="">
				
			
		
		
		
		  <td id="" width="150" valign="">
			
			  
		  	  
		      
				
		  	    <label for="" class="labelClass" style="text-align:right;font-weight:bold;;">Juiz:</label>
		      
		    
		  </td>
		
		
		
		  
		  
		    <td valign="">
		  
		
			
	
	
	
	
	
	<span id="" class="">José Cícero Alves da Silva</span>
	
	
	
	    
	    
    
    
	
	
	
		
		</td>
		
			</tr>
		
	
	





    
    


































































	
	
		
		
	










    
    


































































	
	
		
		
	






	
	
		
		
			
				<tr class="">
				
			
		
		
		
		  <td id="" width="150" valign="">
			
			  
		  	  
		      
				
		  	    <label for="" class="labelClass" style="text-align:right;font-weight:bold;;">Valor da ação:</label>
		      
		    
		  </td>
		
		
		
		  
		  
		    <td valign="">
		  
		
			
	
	
	
	
	
	<span id="" class="">R$         281.178,42</span>
	
	
	
	    
	    
    
    
	
	
	
		
		</td>
		
			</tr>
		
	
	





    
    


































































	
	
		
		
	










    

    
    











































	
	
		
		
	









    
    











































	
	
		
		
	






	
	
		
		
			
				<tr class="">
				
			
		
		
		
		  <td id="" width="150" valign="">
			
			  
		  	  
		      
				
		  	    <label for="tag.dados.rotulo.custas" class="labelClass" style="text-align:right;font-weight:bold;;">Custas:</label>
		      
		    
		  </td>
		
		
		
		  
		  
		    <td valign="">
		  
		
			
	
	
	
	
	
	<table cellpadding="0" cellspacing="0" border="0" width="">
		<tr>
			<td>
				
				
				
        <a href="https://www2.tjal.jus.br/ccpweb/abrirConsultaCustas.do?cdProcesso=01000O7550000&nuProcesso=0710802-55.2018.8.02.0001" target="_blank">Visualizar custas</a>
        
            <span style="color: red; font-weight: bold;">(há custas pendentes)</span>
        
    
				
			</td>
		</tr>
	  </table>
	
	
	
	
	
	
		
		</td>
		
			</tr>
		
	
	




    
    











































	
	
		
		
	









    


		</table>
	</div>
	
	
	



	                
	                
	                
	                
	                
	                
	                
	                
	                
	                
	            
	        
	            
	                
	                
	                    













	                
	                
	                
	                
	                
	                
	                
	                
	                
	            
	        
	            
	                
	                
	                
	                    



























<div style="padding-top: 10px;">
	
	
		
		
		
		
		 
	

	
		
		
			









<table width="100%" border="0" cellspacing="0" cellpadding="0">
	<tr valign="top">
		<td height="21" valign="top" nowrap background="/cpopg/imagens/spw/fundo_subtitulo.gif"> 
			
				
					<h2 class="subtitle">
						Partes do processo
						
					</h2>
				
				
			
		</td>
		<td background="/cpopg/imagens/spw/fundo_subtitulo2.gif" width="90%" aria-hidden="true">
			<img src="/cpopg/imagens/spw/final_subtitulo.gif" width="16" height="20" tabindex="-1">
		</td>
	</tr>
</table>

		
	
</div>


	<div id="divLinksTituloBlocoPartes">
		











<input id="mensagemNaoExibidapartes" type="hidden" value="Exibindo todas as partes." />  
<input id="linkNaoExibidopartes" type="hidden" value="Exibir somente as partes principais." />  
	
<span id="mensagensExibindopartes" class="mensagemExibindo">Exibindo Somente as principais partes.</span> &nbsp; <a id="linkpartes" href="javascript:" class="linkNaoSelecionado" ><span id="setasDireitapartes" class="setasDireita" >&gt;&gt;</span>Exibir todas as partes.</a>

<script>
	$(function() {
		var controlarLink = function() {
			var $linkNaoExibido = $('input#linkNaoExibidopartes');
			var conteudoLinkNaoApresentado = $linkNaoExibido.val();
			var $link = $('a#linkpartes');
			
			$linkNaoExibido.val($link.html());
			$link.html(conteudoLinkNaoApresentado);
		};

		var controlarMensagem = function() {
			var $mensagemNaoExibida = $('input#mensagemNaoExibidapartes');
			var $spanMensagem = $('span#mensagensExibindopartes');
			var mensagemExibida = $spanMensagem.html();
			var mensagemNaoExibida = $mensagemNaoExibida.val();
			
			$spanMensagem.html(mensagemNaoExibida);
			$mensagemNaoExibida.val(mensagemExibida);
		};

		var controlarMensagemLink = function() {
			controlarMensagem();
			controlarLink();
		};
		
		var esconderElementosExtrasMostrarDefault = function() {
			$('#tableTodasPartes').hide();
			$('#tablePartesPrincipais').show();
		};

		var mostrarElementosExtrasEsconderDefault = function() {
			$('#tablePartesPrincipais').hide();
			$('#tableTodasPartes').show();
		};

		var initPagina = function() {
			var setasDireita = '<span id="setasDireitapartes" class="setasDireita">&gt;&gt;</span>';
			var $linkEscondido = $('input#linkNaoExibidopartes');
			$linkEscondido.val(setasDireita+$linkEscondido.val());
		};

		var bindLink = function() {
			var $link = $('a#linkpartes');
			$link.funcToggle('click', mostrarElementosExtrasEsconderDefault, esconderElementosExtrasMostrarDefault);
			$link.bind('click', controlarMensagemLink);
		};

		initPagina();
		bindLink();
		esconderElementosExtrasMostrarDefault();
	});
</script>

	</div>

<!--  cabecalho da tabela de lista (partes) -->


<!--  dados da lista partes principais (partes) -->
<table id="tablePartesPrincipais" style="margin-left:15px; margin-top:1px;" align="center" border="0" cellPadding="0" cellSpacing="0" width="98%">
	
		









<tr class="fundoClaro">
	<td valign="top" align="right" width="141" style="padding-bottom: 5px">
		<span class="mensagemExibindo">Autor:&nbsp;</span>
	</td>
	<td width="*" align="left" style="padding-bottom: 5px">
		
			
				
				
					José Carlos Cerqueira Souza Filho
				
			

			
		  
			<br />
			<span class="mensagemExibindo">Advogado:&nbsp;</span>
			Vinicius Faria de Cerqueira&nbsp;
   	 	
		 
		</td>
</tr>

	
		









<tr class="fundoClaro">
	<td valign="top" align="right" width="141" style="padding-bottom: 5px">
		<span class="mensagemExibindo">Ré:&nbsp;</span>
	</td>
	<td width="*" align="left" style="padding-bottom: 5px">
		
			
				
				
					Cony Engenharia Ltda.
				
			

			
		  
			<br />
			<span class="mensagemExibindo">Advogado:&nbsp;</span>
			Marcus Vinicius Cavalcante Lins Filho&nbsp;
   	 	
		
			<br />
			<span class="mensagemExibindo">Advogado:&nbsp;</span>
			Thiago Maia Nobre Rocha&nbsp;
   	 	
		
			<br />
			<span class="mensagemExibindo">Advogado:&nbsp;</span>
			Orlando de Moura Cavalcante Neto&nbsp;
   	 	
		 
		</td>
</tr>

	
</table>	


	<!--  dados da lista todas as partes (partes) -->
	<table id="tableTodasPartes" style="margin-left:15px; margin-top:1px; display: none; " align="center" width="98%" border="0" cellspacing="0" cellpadding="0"  >
		
			









<tr class="fundoClaro">
	<td valign="top" align="right" width="141" style="padding-bottom: 5px">
		<span class="mensagemExibindo">Autor:&nbsp;</span>
	</td>
	<td width="*" align="left" style="padding-bottom: 5px">
		
			
				
				
					José Carlos Cerqueira Souza Filho
				
			

			
		  
			<br />
			<span class="mensagemExibindo">Advogado:&nbsp;</span>
			Vinicius Faria de Cerqueira&nbsp;
   	 	
		 
		</td>
</tr>

		
			









<tr class="fundoClaro">
	<td valign="top" align="right" width="141" style="padding-bottom: 5px">
		<span class="mensagemExibindo">Autora:&nbsp;</span>
	</td>
	<td width="*" align="left" style="padding-bottom: 5px">
		
			
				
				
					Livia Nascimento da Rocha
				
			

			
		  
			<br />
			<span class="mensagemExibindo">Advogado:&nbsp;</span>
			Vinicius Faria de Cerqueira&nbsp;
   	 	
		 
		</td>
</tr>

		
			









<tr class="fundoClaro">
	<td valign="top" align="right" width="141" style="padding-bottom: 5px">
		<span class="mensagemExibindo">Ré:&nbsp;</span>
	</td>
	<td width="*" align="left" style="padding-bottom: 5px">
		
			
				
				
					Cony Engenharia Ltda.
				
			

			
		  
			<br />
			<span class="mensagemExibindo">Advogado:&nbsp;</span>
			Marcus Vinicius Cavalcante Lins Filho&nbsp;
   	 	
		
			<br />
			<span class="mensagemExibindo">Advogado:&nbsp;</span>
			Thiago Maia Nobre Rocha&nbsp;
   	 	
		
			<br />
			<span class="mensagemExibindo">Advogado:&nbsp;</span>
			Orlando de Moura Cavalcante Neto&nbsp;
   	 	
		 
		</td>
</tr>

		
			









<tr class="fundoClaro">
	<td valign="top" align="right" width="141" style="padding-bottom: 5px">
		<span class="mensagemExibindo">Réu:&nbsp;</span>
	</td>
	<td width="*" align="left" style="padding-bottom: 5px">
		
			
				
				
					Banco do Brasil S A
				
			

			
		  
			<br />
			<span class="mensagemExibindo">Advogado:&nbsp;</span>
			Rafael Sganzerla Durand&nbsp;
   	 	
		 
		</td>
</tr>

		
	</table>


	                
	                
	                
	                
	                
	                
	                
	                
	            
	        
	            
	                
	                
	                
	                
	                
	                
	                    


  






















<div style="padding-top: 10px;">
	
	
		
		
		
		
	

	
		
		
			









<table width="100%" border="0" cellspacing="0" cellpadding="0">
	<tr valign="top">
		<td height="21" valign="top" nowrap background="/cpopg/imagens/spw/fundo_subtitulo.gif"> 
			
				
					<h2 class="subtitle">
						Movimentações
						
					</h2>
				
				
			
		</td>
		<td background="/cpopg/imagens/spw/fundo_subtitulo2.gif" width="90%" aria-hidden="true">
			<img src="/cpopg/imagens/spw/final_subtitulo.gif" width="16" height="20" tabindex="-1">
		</td>
	</tr>
</table>

		
	
</div>


	<div id="divLinksTituloBlocoMovimentacoes">
		











<input id="mensagemNaoExibidamovimentacoes" type="hidden" value="Exibindo todas as movimentações." />  
<input id="linkNaoExibidomovimentacoes" type="hidden" value="Listar somente as 5 últimas." />  
	
<span id="mensagensExibindomovimentacoes" class="mensagemExibindo">Exibindo 5 últimas.</span> &nbsp; <a id="linkmovimentacoes" href="javascript:" class="linkNaoSelecionado" ><span id="setasDireitamovimentacoes" class="setasDireita" >&gt;&gt;</span>Listar todas as movimentações.</a>

<script>
	$(function() {
		var controlarLink = function() {
			var $linkNaoExibido = $('input#linkNaoExibidomovimentacoes');
			var conteudoLinkNaoApresentado = $linkNaoExibido.val();
			var $link = $('a#linkmovimentacoes');
			
			$linkNaoExibido.val($link.html());
			$link.html(conteudoLinkNaoApresentado);
		};

		var controlarMensagem = function() {
			var $mensagemNaoExibida = $('input#mensagemNaoExibidamovimentacoes');
			var $spanMensagem = $('span#mensagensExibindomovimentacoes');
			var mensagemExibida = $spanMensagem.html();
			var mensagemNaoExibida = $mensagemNaoExibida.val();
			
			$spanMensagem.html(mensagemNaoExibida);
			$mensagemNaoExibida.val(mensagemExibida);
		};

		var controlarMensagemLink = function() {
			controlarMensagem();
			controlarLink();
		};
		
		var esconderElementosExtrasMostrarDefault = function() {
			$('#tabelaTodasMovimentacoes').hide();
			$('#tabelaUltimasMovimentacoes').show();
		};

		var mostrarElementosExtrasEsconderDefault = function() {
			$('#tabelaUltimasMovimentacoes').hide();
			$('#tabelaTodasMovimentacoes').show();
		};

		var initPagina = function() {
			var setasDireita = '<span id="setasDireitamovimentacoes" class="setasDireita">&gt;&gt;</span>';
			var $linkEscondido = $('input#linkNaoExibidomovimentacoes');
			$linkEscondido.val(setasDireita+$linkEscondido.val());
		};

		var bindLink = function() {
			var $link = $('a#linkmovimentacoes');
			$link.funcToggle('click', mostrarElementosExtrasEsconderDefault, esconderElementosExtrasMostrarDefault);
			$link.bind('click', controlarMensagemLink);
		};

		initPagina();
		bindLink();
		esconderElementosExtrasMostrarDefault();
	});
</script>

	</div>

 
	

<table  style="margin-left:15px; margin-top:1px;" align="center" border="0" cellPadding="0" cellSpacing="0" width="98%">
	
		
			<thead>
				<tr>
				  <th width="120" class="label" style="vertical-align: bottom">Data</th>
				  <th vAlign="middle" width="20" aria-hidden="true">&nbsp;</th>
				  <th vAlign="middle"  class="label">Movimento</th>
				</tr>
				<tr class="fundoEscuro" height="2" aria-hidden="true">
					<td></td>
					<td></td>
					<td></td>
				</tr>
			</thead>

			
			<tbody id="tabelaUltimasMovimentacoes">
				
					









	
	
		
	


<tr class="fundoClaro" style="">
	<td width="120" style="vertical-align: top">
		12/07/2019
	</td>
	<td width="20" valign="top" aria-hidden="true">
		
	</td>
	<td style="vertical-align: top; padding-bottom: 5px">
		 
		
			
			
				Juntada de Petição
			
		 
		 
		<br /> 
		<span style="font-style: italic;">
			N&ordm; Protocolo: WMAC.19.70150828-9
Tipo da Peti&ccedil;&atilde;o: Peti&ccedil;&atilde;o
Data: 11/07/2019 23:50

		</span> 
		
	</td>
</tr>


				
					









	
	
		
	


<tr class="fundoEscuro" style="">
	<td width="120" style="vertical-align: top">
		12/02/2019
	</td>
	<td width="20" valign="top" aria-hidden="true">
		
	</td>
	<td style="vertical-align: top; padding-bottom: 5px">
		 
		
			
			
				Juntada de Petição
			
		 
		 
		<br /> 
		<span style="font-style: italic;">
			N&ordm; Protocolo: WMAC.19.70034823-7
Tipo da Peti&ccedil;&atilde;o: Peti&ccedil;&atilde;o
Data: 12/02/2019 14:58

		</span> 
		
	</td>
</tr>


				
					









	
	
		
	


<tr class="fundoClaro" style="">
	<td width="120" style="vertical-align: top">
		11/02/2019
	</td>
	<td width="20" valign="top" aria-hidden="true">
		
	</td>
	<td style="vertical-align: top; padding-bottom: 5px">
		 
		
			
			
				Juntada de Petição
			
		 
		 
		<br /> 
		<span style="font-style: italic;">
			N&ordm; Protocolo: WMAC.19.70032532-6
Tipo da Peti&ccedil;&atilde;o: Peti&ccedil;&atilde;o
Data: 11/02/2019 09:13

		</span> 
		
	</td>
</tr>


				
					









	
	
		
	


<tr class="fundoEscuro" style="">
	<td width="120" style="vertical-align: top">
		06/12/2018
	</td>
	<td width="20" valign="top" aria-hidden="true">
		
	</td>
	<td style="vertical-align: top; padding-bottom: 5px">
		 
		
			
			
				Conclusos
			
		 
		 
		<br /> 
		<span style="font-style: italic;">
			
		</span> 
		
	</td>
</tr>


				
					









	
	
		
	


<tr class="fundoClaro" style="">
	<td width="120" style="vertical-align: top">
		05/12/2018
	</td>
	<td width="20" valign="top" aria-hidden="true">
		
	</td>
	<td style="vertical-align: top; padding-bottom: 5px">
		 
		
			
			
				Juntada de Petição
			
		 
		 
		<br /> 
		<span style="font-style: italic;">
			N&ordm; Protocolo: WMAC.18.70259903-1
Tipo da Peti&ccedil;&atilde;o: Peti&ccedil;&atilde;o
Data: 05/12/2018 16:46

		</span> 
		
	</td>
</tr>


				
			</tbody>
			
			
			<tbody style="display: none;" id="tabelaTodasMovimentacoes">
				
					









	
	
		
	


<tr class="fundoClaro" style="">
	<td width="120" style="vertical-align: top">
		12/07/2019
	</td>
	<td width="20" valign="top" aria-hidden="true">
		
	</td>
	<td style="vertical-align: top; padding-bottom: 5px">
		 
		
			
			
				Juntada de Petição
			
		 
		 
		<br /> 
		<span style="font-style: italic;">
			N&ordm; Protocolo: WMAC.19.70150828-9
Tipo da Peti&ccedil;&atilde;o: Peti&ccedil;&atilde;o
Data: 11/07/2019 23:50

		</span> 
		
	</td>
</tr>


				
					









	
	
		
	


<tr class="fundoEscuro" style="">
	<td width="120" style="vertical-align: top">
		12/02/2019
	</td>
	<td width="20" valign="top" aria-hidden="true">
		
	</td>
	<td style="vertical-align: top; padding-bottom: 5px">
		 
		
			
			
				Juntada de Petição
			
		 
		 
		<br /> 
		<span style="font-style: italic;">
			N&ordm; Protocolo: WMAC.19.70034823-7
Tipo da Peti&ccedil;&atilde;o: Peti&ccedil;&atilde;o
Data: 12/02/2019 14:58

		</span> 
		
	</td>
</tr>


				
					









	
	
		
	


<tr class="fundoClaro" style="">
	<td width="120" style="vertical-align: top">
		11/02/2019
	</td>
	<td width="20" valign="top" aria-hidden="true">
		
	</td>
	<td style="vertical-align: top; padding-bottom: 5px">
		 
		
			
			
				Juntada de Petição
			
		 
		 
		<br /> 
		<span style="font-style: italic;">
			N&ordm; Protocolo: WMAC.19.70032532-6
Tipo da Peti&ccedil;&atilde;o: Peti&ccedil;&atilde;o
Data: 11/02/2019 09:13

		</span> 
		
	</td>
</tr>


				
					









	
	
		
	


<tr class="fundoEscuro" style="">
	<td width="120" style="vertical-align: top">
		06/12/2018
	</td>
	<td width="20" valign="top" aria-hidden="true">
		
	</td>
	<td style="vertical-align: top; padding-bottom: 5px">
		 
		
			
			
				Conclusos
			
		 
		 
		<br /> 
		<span style="font-style: italic;">
			
		</span> 
		
	</td>
</tr>


				
					









	
	
		
	


<tr class="fundoClaro" style="">
	<td width="120" style="vertical-align: top">
		05/12/2018
	</td>
	<td width="20" valign="top" aria-hidden="true">
		
	</td>
	<td style="vertical-align: top; padding-bottom: 5px">
		 
		
			
			
				Juntada de Petição
			
		 
		 
		<br /> 
		<span style="font-style: italic;">
			N&ordm; Protocolo: WMAC.18.70259903-1
Tipo da Peti&ccedil;&atilde;o: Peti&ccedil;&atilde;o
Data: 05/12/2018 16:46

		</span> 
		
	</td>
</tr>


				
					









	
	
		
	


<tr class="fundoEscuro" style="">
	<td width="120" style="vertical-align: top">
		29/11/2018
	</td>
	<td width="20" valign="top" aria-hidden="true">
		
	</td>
	<td style="vertical-align: top; padding-bottom: 5px">
		 
		
			
			
				Juntada de Petição
			
		 
		 
		<br /> 
		<span style="font-style: italic;">
			N&ordm; Protocolo: WMAC.18.70255192-6
Tipo da Peti&ccedil;&atilde;o: Peti&ccedil;&atilde;o
Data: 29/11/2018 15:07

		</span> 
		
	</td>
</tr>


				
					









	
	
		
	


<tr class="fundoClaro" style="">
	<td width="120" style="vertical-align: top">
		28/11/2018
	</td>
	<td width="20" valign="top" aria-hidden="true">
		
	</td>
	<td style="vertical-align: top; padding-bottom: 5px">
		 
		
			
			
				Ato Publicado
			
		 
		 
		<br /> 
		<span style="font-style: italic;">
			Rela&ccedil;&atilde;o :0499/2018
Data da Publica&ccedil;&atilde;o: 29/11/2018
N&uacute;mero do Di&aacute;rio: 2233
		</span> 
		
	</td>
</tr>


				
					









	
	
		
	


<tr class="fundoEscuro" style="">
	<td width="120" style="vertical-align: top">
		27/11/2018
	</td>
	<td width="20" valign="top" aria-hidden="true">
		
	</td>
	<td style="vertical-align: top; padding-bottom: 5px">
		 
		
			
			
				Encaminhado ao DJ Eletrônico
			
		 
		 
		<br /> 
		<span style="font-style: italic;">
			Rela&ccedil;&atilde;o: 0499/2018
Teor do ato: ATO ORDINAT&Oacute;RIO Em cumprimento ao disposto no art. 152,VI do CPC, procedo &agrave; intima&ccedil;&atilde;o das partes para especificarem e justificarem as provas que ainda pretendem produzir, fundamentamente, pelo prazo comum de 5(cinco) dias. Macei&oacute;, 27 de novembro de 2018
Advogados(s): Orlando de Moura Cavalcante Neto (OAB 7313/AL), Thiago Maia Nobre Rocha (OAB 6213/AL), Vinicius Faria de Cerqueira (OAB 9008/AL), Rafael Sganzerla Durand (OAB 10132A/AL), Marcus Vinicius Cavalcante Lins Filho (OAB 10871/AL)
		</span> 
		
	</td>
</tr>


				
					









	
	
		
	


<tr class="fundoClaro" style="">
	<td width="120" style="vertical-align: top">
		27/11/2018
	</td>
	<td width="20" valign="top" aria-hidden="true">
		
			
				
				
					
				
			 
						
			<a class="linkMovVincProc"
				id="linkMovVincProc-26689702"
				title="Visualizar documento em inteiro teor"
				href="#liberarAutoPorSenha">
				<img width="16" height="16" border="0" src="/cpopg/imagens/doc2.gif" />
			</a>
		
	</td>
	<td style="vertical-align: top; padding-bottom: 5px">
		 
		
			
				<a class="linkMovVincProc"
					id="linkMovVincProc-2-26689702"
					title="Visualizar documento em inteiro teor"
					href="#liberarAutoPorSenha"> Ato ordinatório praticado
				</a>
			
			
		 
		 
		<br /> 
		<span style="font-style: italic;">
			ATO ORDINAT&Oacute;RIO Em cumprimento ao disposto no art. 152,VI do CPC, procedo &agrave; intima&ccedil;&atilde;o das partes para especificarem e justificarem as provas que ainda pretendem produzir, fundamentamente, pelo prazo comum de 5(cinco) dias. Macei&oacute;, 27 de novembro de 2018
		</span> 
		
	</td>
</tr>


				
					









	
	
		
	


<tr class="fundoEscuro" style="">
	<td width="120" style="vertical-align: top">
		26/11/2018
	</td>
	<td width="20" valign="top" aria-hidden="true">
		
	</td>
	<td style="vertical-align: top; padding-bottom: 5px">
		 
		
			
			
				Juntada de Documento
			
		 
		 
		<br /> 
		<span style="font-style: italic;">
			N&ordm; Protocolo: WMAC.18.70251514-8
Tipo da Peti&ccedil;&atilde;o: Impugna&ccedil;&atilde;o &agrave; Contesta&ccedil;&atilde;o
Data: 26/11/2018 15:37

		</span> 
		
	</td>
</tr>


				
					









	
	
		
	


<tr class="fundoClaro" style="">
	<td width="120" style="vertical-align: top">
		26/11/2018
	</td>
	<td width="20" valign="top" aria-hidden="true">
		
	</td>
	<td style="vertical-align: top; padding-bottom: 5px">
		 
		
			
			
				Juntada de Documento
			
		 
		 
		<br /> 
		<span style="font-style: italic;">
			N&ordm; Protocolo: WMAC.18.70251511-3
Tipo da Peti&ccedil;&atilde;o: Impugna&ccedil;&atilde;o &agrave; Contesta&ccedil;&atilde;o
Data: 26/11/2018 15:35

		</span> 
		
	</td>
</tr>


				
					









	
	
		
	


<tr class="fundoEscuro" style="">
	<td width="120" style="vertical-align: top">
		09/11/2018
	</td>
	<td width="20" valign="top" aria-hidden="true">
		
	</td>
	<td style="vertical-align: top; padding-bottom: 5px">
		 
		
			
			
				Ato Publicado
			
		 
		 
		<br /> 
		<span style="font-style: italic;">
			Rela&ccedil;&atilde;o :0456/2018
Data da Publica&ccedil;&atilde;o: 12/11/2018
N&uacute;mero do Di&aacute;rio: 2222
		</span> 
		
	</td>
</tr>


				
					









	
	
		
	


<tr class="fundoClaro" style="">
	<td width="120" style="vertical-align: top">
		08/11/2018
	</td>
	<td width="20" valign="top" aria-hidden="true">
		
	</td>
	<td style="vertical-align: top; padding-bottom: 5px">
		 
		
			
			
				Encaminhado ao DJ Eletrônico
			
		 
		 
		<br /> 
		<span style="font-style: italic;">
			Rela&ccedil;&atilde;o: 0456/2018
Teor do ato: Autos n&deg;: 0710802-55.2018.8.02.0001 A&ccedil;&atilde;o: Procedimento Ordin&aacute;rio Autor: Jos&eacute; Carlos Cerqueira Souza Filho e outro R&eacute;u: Cona&ccedil;o Engenharia Ltda. e outro ATO ORDINAT&Oacute;RIO Intimo a parte autora a apresentar, querendo, no prazo de 15 (quinze) dias, impugna&ccedil;&atilde;o &agrave;s contesta&ccedil;&otilde;es. Macei&oacute;, 06 de novembro de 2018 Hallph S&aacute; de Ara&uacute;jo Analista Judici&aacute;rio
Advogados(s): Vinicius Faria de Cerqueira (OAB 9008/AL)
		</span> 
		
	</td>
</tr>


				
					









	
	
		
	


<tr class="fundoEscuro" style="">
	<td width="120" style="vertical-align: top">
		06/11/2018
	</td>
	<td width="20" valign="top" aria-hidden="true">
		
			
				
				
					
				
			 
						
			<a class="linkMovVincProc"
				id="linkMovVincProc-26414651"
				title="Visualizar documento em inteiro teor"
				href="#liberarAutoPorSenha">
				<img width="16" height="16" border="0" src="/cpopg/imagens/doc2.gif" />
			</a>
		
	</td>
	<td style="vertical-align: top; padding-bottom: 5px">
		 
		
			
				<a class="linkMovVincProc"
					id="linkMovVincProc-2-26414651"
					title="Visualizar documento em inteiro teor"
					href="#liberarAutoPorSenha"> Ato ordinatório praticado
				</a>
			
			
		 
		 
		<br /> 
		<span style="font-style: italic;">
			Autos n&deg;: 0710802-55.2018.8.02.0001 A&ccedil;&atilde;o: Procedimento Ordin&aacute;rio Autor: Jos&eacute; Carlos Cerqueira Souza Filho e outro R&eacute;u: Cona&ccedil;o Engenharia Ltda. e outro ATO ORDINAT&Oacute;RIO Intimo a parte autora a apresentar, querendo, no prazo de 15 (quinze) dias, impugna&ccedil;&atilde;o &agrave;s contesta&ccedil;&otilde;es. Macei&oacute;, 06 de novembro de 2018 Hallph S&aacute; de Ara&uacute;jo Analista Judici&aacute;rio<br/><b>Vencimento: </b>29/11/2018
		</span> 
		
	</td>
</tr>


				
					









	
	
		
	


<tr class="fundoClaro" style="">
	<td width="120" style="vertical-align: top">
		16/10/2018
	</td>
	<td width="20" valign="top" aria-hidden="true">
		
	</td>
	<td style="vertical-align: top; padding-bottom: 5px">
		 
		
			
			
				Juntada de Documentos
			
		 
		 
		<br /> 
		<span style="font-style: italic;">
			N&ordm; Protocolo: WMAC.18.70221617-5
Tipo da Peti&ccedil;&atilde;o: Contesta&ccedil;&atilde;o
Data: 16/10/2018 14:43

		</span> 
		
	</td>
</tr>


				
					









	
	
		
	


<tr class="fundoEscuro" style="">
	<td width="120" style="vertical-align: top">
		10/10/2018
	</td>
	<td width="20" valign="top" aria-hidden="true">
		
	</td>
	<td style="vertical-align: top; padding-bottom: 5px">
		 
		
			
			
				Juntada de Documentos
			
		 
		 
		<br /> 
		<span style="font-style: italic;">
			N&ordm; Protocolo: WMAC.18.70218154-1
Tipo da Peti&ccedil;&atilde;o: Contesta&ccedil;&atilde;o
Data: 10/10/2018 14:04

		</span> 
		
	</td>
</tr>


				
					









	
	
		
	


<tr class="fundoClaro" style="">
	<td width="120" style="vertical-align: top">
		24/09/2018
	</td>
	<td width="20" valign="top" aria-hidden="true">
		
	</td>
	<td style="vertical-align: top; padding-bottom: 5px">
		 
		
			
			
				Juntada de Documento
			
		 
		 
		<br /> 
		<span style="font-style: italic;">
			
		</span> 
		
	</td>
</tr>


				
					









	
	
		
	


<tr class="fundoEscuro" style="">
	<td width="120" style="vertical-align: top">
		24/09/2018
	</td>
	<td width="20" valign="top" aria-hidden="true">
		
	</td>
	<td style="vertical-align: top; padding-bottom: 5px">
		 
		
			
			
				Juntada de Documento
			
		 
		 
		<br /> 
		<span style="font-style: italic;">
			
		</span> 
		
	</td>
</tr>


				
					









	
	
		
	


<tr class="fundoClaro" style="">
	<td width="120" style="vertical-align: top">
		24/09/2018
	</td>
	<td width="20" valign="top" aria-hidden="true">
		
			
				
				
					
				
			 
						
			<a class="linkMovVincProc"
				id="linkMovVincProc-25822019"
				title="Visualizar documento em inteiro teor"
				href="#liberarAutoPorSenha">
				<img width="16" height="16" border="0" src="/cpopg/imagens/doc2.gif" />
			</a>
		
	</td>
	<td style="vertical-align: top; padding-bottom: 5px">
		 
		
			
				<a class="linkMovVincProc"
					id="linkMovVincProc-2-25822019"
					title="Visualizar documento em inteiro teor"
					href="#liberarAutoPorSenha"> Audiência Realizada
				</a>
			
			
		 
		 
		<br /> 
		<span style="font-style: italic;">
			Assentada - Gen&eacute;rico
		</span> 
		
	</td>
</tr>


				
					









	
	
		
	


<tr class="fundoEscuro" style="">
	<td width="120" style="vertical-align: top">
		24/09/2018
	</td>
	<td width="20" valign="top" aria-hidden="true">
		
	</td>
	<td style="vertical-align: top; padding-bottom: 5px">
		 
		
			
			
				Juntada de Petição
			
		 
		 
		<br /> 
		<span style="font-style: italic;">
			N&ordm; Protocolo: WMAC.18.70203989-3
Tipo da Peti&ccedil;&atilde;o: Peti&ccedil;&atilde;o
Data: 24/09/2018 10:09

		</span> 
		
	</td>
</tr>


				
					









	
	
		
	


<tr class="fundoClaro" style="">
	<td width="120" style="vertical-align: top">
		21/09/2018
	</td>
	<td width="20" valign="top" aria-hidden="true">
		
	</td>
	<td style="vertical-align: top; padding-bottom: 5px">
		 
		
			
			
				Juntada de Petição
			
		 
		 
		<br /> 
		<span style="font-style: italic;">
			N&ordm; Protocolo: WMAC.18.70203544-8
Tipo da Peti&ccedil;&atilde;o: Peti&ccedil;&atilde;o
Data: 21/09/2018 18:07

		</span> 
		
	</td>
</tr>


				
					









	
	
		
	


<tr class="fundoEscuro" style="">
	<td width="120" style="vertical-align: top">
		21/09/2018
	</td>
	<td width="20" valign="top" aria-hidden="true">
		
	</td>
	<td style="vertical-align: top; padding-bottom: 5px">
		 
		
			
			
				Juntada de Petição
			
		 
		 
		<br /> 
		<span style="font-style: italic;">
			N&ordm; Protocolo: WMAC.18.70203528-6
Tipo da Peti&ccedil;&atilde;o: Peti&ccedil;&atilde;o
Data: 21/09/2018 17:42

		</span> 
		
	</td>
</tr>


				
					









	
	
		
	


<tr class="fundoClaro" style="">
	<td width="120" style="vertical-align: top">
		21/09/2018
	</td>
	<td width="20" valign="top" aria-hidden="true">
		
	</td>
	<td style="vertical-align: top; padding-bottom: 5px">
		 
		
			
			
				Juntada de Petição
			
		 
		 
		<br /> 
		<span style="font-style: italic;">
			N&ordm; Protocolo: WMAC.18.70203260-0
Tipo da Peti&ccedil;&atilde;o: Peti&ccedil;&atilde;o
Data: 21/09/2018 13:58

		</span> 
		
	</td>
</tr>


				
					









	
		
	
	


<tr class="fundoEscuro" style="">
	<td width="120" style="vertical-align: top">
		20/09/2018
	</td>
	<td width="20" valign="top" aria-hidden="true">
		
			
				
				
					
				
			 
						
			<a class="linkMovVincProc"
				id="linkMovVincProc-25787004"
				title="Visualizar documento em inteiro teor"
				href="/cpopg/abrirDocumentoVinculadoMovimentacao.do?processo.codigo=01000O7550000&cdDocumento=25787004&nmRecursoAcessado=Visto+em+correi%C3%A7%C3%A3o">
				<img width="16" height="16" border="0" src="/cpopg/imagens/doc2.gif" />
			</a>
		
	</td>
	<td style="vertical-align: top; padding-bottom: 5px">
		 
		
			
				<a class="linkMovVincProc"
					id="linkMovVincProc-2-25787004"
					title="Visualizar documento em inteiro teor"
					href="/cpopg/abrirDocumentoVinculadoMovimentacao.do?processo.codigo=01000O7550000&cdDocumento=25787004&nmRecursoAcessado=Visto+em+correi%C3%A7%C3%A3o"> Visto em correição
				</a>
			
			
		 
		 
		<br /> 
		<span style="font-style: italic;">
			DESPACHO VISTO EM CORREI&Ccedil;&Atilde;O
		</span> 
		
	</td>
</tr>


				
					









	
	
		
	


<tr class="fundoClaro" style="">
	<td width="120" style="vertical-align: top">
		06/06/2018
	</td>
	<td width="20" valign="top" aria-hidden="true">
		
	</td>
	<td style="vertical-align: top; padding-bottom: 5px">
		 
		
			
			
				Juntada de AR - Cumprido
			
		 
		 
		<br /> 
		<span style="font-style: italic;">
			Em 06  de  junho  de  2018 &eacute; juntado a estes autos o aviso de recebimento (AR803969759TJ - Cumprido), referente  ao  of&iacute;cio  n. 0710802-55.2018.8.02.0001-0002, emitido para Cona&ccedil;o Engenharia Ltda.. Usu&aacute;rio: 
		</span> 
		
	</td>
</tr>


				
					









	
	
		
	


<tr class="fundoEscuro" style="">
	<td width="120" style="vertical-align: top">
		06/06/2018
	</td>
	<td width="20" valign="top" aria-hidden="true">
		
	</td>
	<td style="vertical-align: top; padding-bottom: 5px">
		 
		
			
			
				Juntada de AR - Cumprido
			
		 
		 
		<br /> 
		<span style="font-style: italic;">
			Em 06  de  junho  de  2018 &eacute; juntado a estes autos o aviso de recebimento (AR803969745TJ - Cumprido), referente  ao  of&iacute;cio  n. 0710802-55.2018.8.02.0001-0001, emitido para Banco do Brasil S A. Usu&aacute;rio: 
		</span> 
		
	</td>
</tr>


				
					









	
	
		
	


<tr class="fundoClaro" style="">
	<td width="120" style="vertical-align: top">
		15/05/2018
	</td>
	<td width="20" valign="top" aria-hidden="true">
		
	</td>
	<td style="vertical-align: top; padding-bottom: 5px">
		 
		
			
			
				Ato Publicado
			
		 
		 
		<br /> 
		<span style="font-style: italic;">
			Rela&ccedil;&atilde;o :0220/2018
Data da Publica&ccedil;&atilde;o: 16/05/2018
N&uacute;mero do Di&aacute;rio: 2105
		</span> 
		
	</td>
</tr>


				
					









	
	
		
	


<tr class="fundoEscuro" style="">
	<td width="120" style="vertical-align: top">
		15/05/2018
	</td>
	<td width="20" valign="top" aria-hidden="true">
		
	</td>
	<td style="vertical-align: top; padding-bottom: 5px">
		 
		
			
			
				Ato Publicado
			
		 
		 
		<br /> 
		<span style="font-style: italic;">
			Rela&ccedil;&atilde;o :0220/2018
Data da Publica&ccedil;&atilde;o: 16/05/2018
N&uacute;mero do Di&aacute;rio: 2105
		</span> 
		
	</td>
</tr>


				
					









	
	
		
	


<tr class="fundoClaro" style="">
	<td width="120" style="vertical-align: top">
		11/05/2018
	</td>
	<td width="20" valign="top" aria-hidden="true">
		
	</td>
	<td style="vertical-align: top; padding-bottom: 5px">
		 
		
			
			
				Encaminhado ao DJ Eletrônico
			
		 
		 
		<br /> 
		<span style="font-style: italic;">
			Rela&ccedil;&atilde;o: 0220/2018
Teor do ato: DECIS&Atilde;OTrata-se de a&ccedil;&atilde;o ordin&aacute;ria de indeniza&ccedil;&atilde;o por danos morais e materias c/c obriga&ccedil;&atilde;o de fazer c/c declara&ccedil;&atilde;o de nulidade de contrato de financiamento banc&aacute;rio c/c pedido de tutela antecipada proposta por JOS&Eacute; CARLOS CERQUEIRA SOUZA FILHO e LIVIA NASCIMENTO DA ROCHA, qualificados na inicial, em face de a CONY ENGENHARIA LTDA. e BANCO DO BRASIL, igualmente qualificados.Narra a exordial que os autores adquiriram o apartamento residencial de n&ordm; 502, da Torre I, do Empreendimento Residencial Dellavia Park Club, situado na Ladeira Murilo Monteiro Valente, n.&ordm; 375, no bairro do Barro Duro, Macei&oacute;/AL, cuja vendedora foi a r&eacute; CONY ENGENHARIA LTDA., pelo valor de R$ 232.000,00 (duzentos e trinta e dois mil reais).Segue narrando que o instrumento celebrado em 02/12/2013 previa a entrega do im&oacute;vel no prazo de 36 (trinta e seis) meses contados do in&iacute;cio da obra, sendo admitida uma toler&acirc;ncia de no m&aacute;ximo 18 (dezoito) meses. Dentro dessa perspectiva, foi prometido que a obra seria iniciada em no m&aacute;ximo 60 (sessenta) dias da assinatura do contrato, ou seja seria iniciada em 02/02/2014 com previs&atilde;o de entrega para 02/02/2017, por&eacute;m, at&eacute; a presente data a obra n&atilde;o foi conclu&iacute;da.Aduz que os demandantes ainda sendo cobrados ilegalmente pelo BANCO DO BRASIL, tamb&eacute;m r&eacute;u da a&ccedil;&atilde;o, numa suposta &quot;taxa de obra&quot;, que decorre de financiamento banc&aacute;rio.Requer, em sede de tutela de urg&ecirc;ncia, que seja determinado ao BANCO DO BRASIL a SUSPENS&Atilde;O da cobran&ccedil;a de taxa de Obra.&Eacute; o relat&oacute;rio. Decido.Ab initio, concedo aos Demandantes as benesses da assist&ecirc;ncia judici&aacute;ria gratuita, em respeito as determina&ccedil;&otilde;es contidas no art. 98 e art. 99 da Lei n&ordm;. 13.105/2015 (C&oacute;digo de Processo Civil - CPC/2015).O C&oacute;digo de Defesa do Consumidor, em seu art. 6.&ordm;, VIII, assegura como direito b&aacute;sico do consumidor a facilita&ccedil;&atilde;o da defesa de seus direitos, inclusive com a invers&atilde;o do &ocirc;nus da prova, a seu favor, quando, a crit&eacute;rio do juiz, for veross&iacute;mil a alega&ccedil;&atilde;o ou quando for ele hipossuficiente, segundo as regras ordin&aacute;rias de experi&ecirc;ncia. Busca-se, assim assegurar a igualdade material.Em que pese bastar apenas um dos requisitos para a invers&atilde;o, o caso em tela preenche as duas condi&ccedil;&otilde;es. Assim com fulcro no art. 6.&ordm;, VIII, do CDC DECIDO POR INVERTER O &Ocirc;NUS DA PROVA.Passo a apreciar o pedido de tutela provis&oacute;ria de urg&ecirc;ncia.Segundo o art. 300 do CPC/15, a tutela de urg&ecirc;ncia ser&aacute; concedida quando houver elementos que evidenciem a probabilidade do direito e o perigo de dano ou o risco ao resultado &uacute;til do processo. O dispositivo deixa evidentes os requisitos da tutela antecipada de urg&ecirc;ncia, quais sejam, a probabilidade do direito, doutrinariamente conhecida como fumus boni iuris, e o perigo de dano ou risco ao resultado &uacute;til do processo, chamado periculum in mora.Nesse trilhar, importa esclarecer que a tutela de urg&ecirc;ncia antecipada se funda em um Ju&iacute;zo de cogni&ccedil;&atilde;o sum&aacute;ria, de modo que a medida, quando concedida, ser&aacute; prec&aacute;ria, haja vista ser fundamental a necessidade de ser revers&iacute;vel (300, &sect;3&ordm;, do CPC/2015).Portanto, a antecipa&ccedil;&atilde;o provis&oacute;ria dos efeitos finais da tutela definitiva, permite o gozo antecipado e imediato dos efeitos pr&oacute;prios da tutela definitiva pretendida, mas n&atilde;o se funda em um ju&iacute;zo de valor exauriente, de modo que pode ser desconstitu&iacute;da a qualquer tempo.Nessa esteira de pensamento, passa-se a analisar o caso concreto e o preenchimento dos requisitos necess&aacute;rios &agrave; concess&atilde;o da tutela provis&oacute;ria pretendida.No caso em tela, pretende a parte autora a suspens&atilde;o da cobran&ccedil;a da &quot;taxa de obra&quot;, em raz&atilde;o do suposto descumprimento contratual por parte da demandada no tocante ao prazo estabelecido para a entrega do im&oacute;vel.Conforme se extrai dos autos, as partes firmaram um contrato de compra e venda de um apartamento, tendo sido estipulada o prazo para sua entrega em 36 (trinta e seis) meses, com um prazo de toler&acirc;ncia de 18 (dezoito) meses, consoante cl&aacute;usula quarta do contrato (fls.39). Logo o prazo final para entrega do im&oacute;vel se encerra em 02/08/2018, levando em considera&ccedil;&atilde;o o prazo de toler&acirc;ncia estabelecido no contrato.&Eacute; cedi&ccedil;o que a taxa de evolu&ccedil;&atilde;o de obra &eacute; devida desde a aprova&ccedil;&atilde;o do financiamento at&eacute; o t&eacute;rmino da obra.&nbsp;Portanto, se a obra atrasar, &eacute; devido o pagamento da referida taxa ao banco que financiou o im&oacute;vel, no caso, o Banco do Brasil, at&eacute; a sua conclus&atilde;o. Sendo certo que ocorrendo a mora da construtora requerida em rela&ccedil;&atilde;o &agrave; entrega do im&oacute;vel, a parte autora n&atilde;o pode ser penalizada com o pagamento de tal encargo.&nbsp;No entanto, no caso em deslinde, ainda n&atilde;o houve mora da construtora, haja vista que ainda n&atilde;o fora encerrado o prazo estimado para entrega do im&oacute;vel. Nesse ponto impende destacar que &eacute; legal a cl&aacute;usula contratual que prev&ecirc; a prorroga&ccedil;&atilde;o do prazo razo&aacute;vel para entrega do im&oacute;vel, considerando o princ&iacute;pio pacta sunt servanda.&nbsp;Desta feita, verifica-se a aus&ecirc;ncia de probabilidade do direito da parte autora, haja vista que, consoante dito alhures, a taxa de evolu&ccedil;&atilde;o de obra &eacute; devida at&eacute; a conclus&atilde;o da obra, bem como que n&atilde;o houve mora por parte da construtora demandada, haja vista que n&atilde;o houve, ainda, atraso na entrega do im&oacute;vel, tendo em vista a cl&aacute;usula que prev&ecirc; prazo de toler&acirc;ncia para entrega do im&oacute;vel.Ante o exposto, por considerar ausente a probabilidade do direito (art. 300 do CPC/15), INDEFIRO o pedido de tutela de urg&ecirc;ncia requestado.Inclua-se o feito em pauta de audi&ecirc;ncia de concilia&ccedil;&atilde;o. Cite-se a parte r&eacute; eintime-a desta decis&atilde;o, bem como para que compare&ccedil;a &agrave; audi&ecirc;ncia na data designada pelo Cart&oacute;rio, o que deve ser feito com anteced&ecirc;ncia m&iacute;nima de 20 dias.Intime-se os autores por advogado constitu&iacute;do (art. 334, &sect;3&ordm;, CPC/15).Dever&aacute; a parte r&eacute; ser advertida da possibilidade do art. 334, &sect;5&ordm;, bem como do termo inicial do prazo de contesta&ccedil;&atilde;o (art. 335).Fiquem as partes advertidas, ainda, de que o n&atilde;o comparecimento injustificado &agrave; audi&ecirc;ncia de concilia&ccedil;&atilde;o &eacute; considerado ato atentat&oacute;rio &agrave; dignidade da justi&ccedil;a e ser&aacute; sancionado com multa de at&eacute; dois por cento da vantagem econ&ocirc;mica pretendida ou do valor da causa (art. 334, &sect;8&ordm;).Publique-se. Intime-seMacei&oacute;, 10 de maio de 2018.Henrique Gomes de Barros TeixeiraJuiz de Direito
Advogados(s): Vinicius Faria de Cerqueira (OAB 9008/AL)
		</span> 
		
	</td>
</tr>


				
					









	
	
		
	


<tr class="fundoEscuro" style="">
	<td width="120" style="vertical-align: top">
		11/05/2018
	</td>
	<td width="20" valign="top" aria-hidden="true">
		
	</td>
	<td style="vertical-align: top; padding-bottom: 5px">
		 
		
			
			
				Encaminhado ao DJ Eletrônico
			
		 
		 
		<br /> 
		<span style="font-style: italic;">
			Rela&ccedil;&atilde;o: 0220/2018
Teor do ato: Concilia&ccedil;&atilde;o
Data: 24/09/2018 Hora 14:00
Local: Sala de Audi&ecirc;ncia
Situac&atilde;o: Pendente
Advogados(s): Vinicius Faria de Cerqueira (OAB 9008/AL)
		</span> 
		
	</td>
</tr>


				
					









	
	
		
	


<tr class="fundoClaro" style="">
	<td width="120" style="vertical-align: top">
		11/05/2018
	</td>
	<td width="20" valign="top" aria-hidden="true">
		
			
				
				
					
				
			 
						
			<a class="linkMovVincProc"
				id="linkMovVincProc-24228218"
				title="Visualizar documento em inteiro teor"
				href="#liberarAutoPorSenha">
				<img width="16" height="16" border="0" src="/cpopg/imagens/doc2.gif" />
			</a>
		
	</td>
	<td style="vertical-align: top; padding-bottom: 5px">
		 
		
			
				<a class="linkMovVincProc"
					id="linkMovVincProc-2-24228218"
					title="Visualizar documento em inteiro teor"
					href="#liberarAutoPorSenha"> Carta Expedida
				</a>
			
			
		 
		 
		<br /> 
		<span style="font-style: italic;">
			AR DIGITAL - Cita&ccedil;&atilde;o e Intima&ccedil;&atilde;o - Audi&ecirc;ncia de Instru&ccedil;&atilde;o e Julgamento - Juizado
		</span> 
		
	</td>
</tr>


				
					









	
	
		
	


<tr class="fundoEscuro" style="">
	<td width="120" style="vertical-align: top">
		11/05/2018
	</td>
	<td width="20" valign="top" aria-hidden="true">
		
			
				
				
					
				
			 
						
			<a class="linkMovVincProc"
				id="linkMovVincProc-24228179"
				title="Visualizar documento em inteiro teor"
				href="#liberarAutoPorSenha">
				<img width="16" height="16" border="0" src="/cpopg/imagens/doc2.gif" />
			</a>
		
	</td>
	<td style="vertical-align: top; padding-bottom: 5px">
		 
		
			
				<a class="linkMovVincProc"
					id="linkMovVincProc-2-24228179"
					title="Visualizar documento em inteiro teor"
					href="#liberarAutoPorSenha"> Carta Expedida
				</a>
			
			
		 
		 
		<br /> 
		<span style="font-style: italic;">
			AR DIGITAL - Cita&ccedil;&atilde;o e Intima&ccedil;&atilde;o - Audi&ecirc;ncia de Instru&ccedil;&atilde;o e Julgamento - Juizado
		</span> 
		
	</td>
</tr>


				
					









	
	
		
	


<tr class="fundoClaro" style="">
	<td width="120" style="vertical-align: top">
		11/05/2018
	</td>
	<td width="20" valign="top" aria-hidden="true">
		
	</td>
	<td style="vertical-align: top; padding-bottom: 5px">
		 
		
			
			
				Audiência Designada
			
		 
		 
		<br /> 
		<span style="font-style: italic;">
			Concilia&ccedil;&atilde;o
Data: 24/09/2018 Hora 14:00
Local: Sala de Audi&ecirc;ncia
Situac&atilde;o: Realizada
		</span> 
		
	</td>
</tr>


				
					









	
		
	
	


<tr class="fundoEscuro" style="">
	<td width="120" style="vertical-align: top">
		10/05/2018
	</td>
	<td width="20" valign="top" aria-hidden="true">
		
			
				
				
					
				
			 
						
			<a class="linkMovVincProc"
				id="linkMovVincProc-24187801"
				title="Visualizar documento em inteiro teor"
				href="/cpopg/abrirDocumentoVinculadoMovimentacao.do?processo.codigo=01000O7550000&cdDocumento=24187801&nmRecursoAcessado=N%C3%A3o+Concedida+a+Antecipa%C3%A7%C3%A3o+de+tutela">
				<img width="16" height="16" border="0" src="/cpopg/imagens/doc2.gif" />
			</a>
		
	</td>
	<td style="vertical-align: top; padding-bottom: 5px">
		 
		
			
				<a class="linkMovVincProc"
					id="linkMovVincProc-2-24187801"
					title="Visualizar documento em inteiro teor"
					href="/cpopg/abrirDocumentoVinculadoMovimentacao.do?processo.codigo=01000O7550000&cdDocumento=24187801&nmRecursoAcessado=N%C3%A3o+Concedida+a+Antecipa%C3%A7%C3%A3o+de+tutela"> Não Concedida a Antecipação de tutela
				</a>
			
			
		 
		 
		<br /> 
		<span style="font-style: italic;">
			DECIS&Atilde;OTrata-se de a&ccedil;&atilde;o ordin&aacute;ria de indeniza&ccedil;&atilde;o por danos morais e materias c/c obriga&ccedil;&atilde;o de fazer c/c declara&ccedil;&atilde;o de nulidade de contrato de financiamento banc&aacute;rio c/c pedido de tutela antecipada proposta por JOS&Eacute; CARLOS CERQUEIRA SOUZA FILHO e LIVIA NASCIMENTO DA ROCHA, qualificados na inicial, em face de a CONY ENGENHARIA LTDA. e BANCO DO BRASIL, igualmente qualificados.Narra a exordial que os autores adquiriram o apartamento residencial de n&ordm; 502, da Torre I, do Empreendimento Residencial Dellavia Park Club, situado na Ladeira Murilo Monteiro Valente, n.&ordm; 375, no bairro do Barro Duro, Macei&oacute;/AL, cuja vendedora foi a r&eacute; CONY ENGENHARIA LTDA., pelo valor de R$ 232.000,00 (duzentos e trinta e dois mil reais).Segue narrando que o instrumento celebrado em 02/12/2013 previa a entrega do im&oacute;vel no prazo de 36 (trinta e seis) meses contados do in&iacute;cio da obra, sendo admitida uma toler&acirc;ncia de no m&aacute;ximo 18 (dezoito) meses. Dentro dessa perspectiva, foi prometido que a obra seria iniciada em no m&aacute;ximo 60 (sessenta) dias da assinatura do contrato, ou seja seria iniciada em 02/02/2014 com previs&atilde;o de entrega para 02/02/2017, por&eacute;m, at&eacute; a presente data a obra n&atilde;o foi conclu&iacute;da.Aduz que os demandantes ainda sendo cobrados ilegalmente pelo BANCO DO BRASIL, tamb&eacute;m r&eacute;u da a&ccedil;&atilde;o, numa suposta &quot;taxa de obra&quot;, que decorre de financiamento banc&aacute;rio.Requer, em sede de tutela de urg&ecirc;ncia, que seja determinado ao BANCO DO BRASIL a SUSPENS&Atilde;O da cobran&ccedil;a de taxa de Obra.&Eacute; o relat&oacute;rio. Decido.Ab initio, concedo aos Demandantes as benesses da assist&ecirc;ncia judici&aacute;ria gratuita, em respeito as determina&ccedil;&otilde;es contidas no art. 98 e art. 99 da Lei n&ordm;. 13.105/2015 (C&oacute;digo de Processo Civil - CPC/2015).O C&oacute;digo de Defesa do Consumidor, em seu art. 6.&ordm;, VIII, assegura como direito b&aacute;sico do consumidor a facilita&ccedil;&atilde;o da defesa de seus direitos, inclusive com a invers&atilde;o do &ocirc;nus da prova, a seu favor, quando, a crit&eacute;rio do juiz, for veross&iacute;mil a alega&ccedil;&atilde;o ou quando for ele hipossuficiente, segundo as regras ordin&aacute;rias de experi&ecirc;ncia. Busca-se, assim assegurar a igualdade material.Em que pese bastar apenas um dos requisitos para a invers&atilde;o, o caso em tela preenche as duas condi&ccedil;&otilde;es. Assim com fulcro no art. 6.&ordm;, VIII, do CDC DECIDO POR INVERTER O &Ocirc;NUS DA PROVA.Passo a apreciar o pedido de tutela provis&oacute;ria de urg&ecirc;ncia.Segundo o art. 300 do CPC/15, a tutela de urg&ecirc;ncia ser&aacute; concedida quando houver elementos que evidenciem a probabilidade do direito e o perigo de dano ou o risco ao resultado &uacute;til do processo. O dispositivo deixa evidentes os requisitos da tutela antecipada de urg&ecirc;ncia, quais sejam, a probabilidade do direito, doutrinariamente conhecida como fumus boni iuris, e o perigo de dano ou risco ao resultado &uacute;til do processo, chamado periculum in mora.Nesse trilhar, importa esclarecer que a tutela de urg&ecirc;ncia antecipada se funda em um Ju&iacute;zo de cogni&ccedil;&atilde;o sum&aacute;ria, de modo que a medida, quando concedida, ser&aacute; prec&aacute;ria, haja vista ser fundamental a necessidade de ser revers&iacute;vel (300, &sect;3&ordm;, do CPC/2015).Portanto, a antecipa&ccedil;&atilde;o provis&oacute;ria dos efeitos finais da tutela definitiva, permite o gozo antecipado e imediato dos efeitos pr&oacute;prios da tutela definitiva pretendida, mas n&atilde;o se funda em um ju&iacute;zo de valor exauriente, de modo que pode ser desconstitu&iacute;da a qualquer tempo.Nessa esteira de pensamento, passa-se a analisar o caso concreto e o preenchimento dos requisitos necess&aacute;rios &agrave; concess&atilde;o da tutela provis&oacute;ria pretendida.No caso em tela, pretende a parte autora a suspens&atilde;o da cobran&ccedil;a da &quot;taxa de obra&quot;, em raz&atilde;o do suposto descumprimento contratual por parte da demandada no tocante ao prazo estabelecido para a entrega do im&oacute;vel.Conforme se extrai dos autos, as partes firmaram um contrato de compra e venda de um apartamento, tendo sido estipulada o prazo para sua entrega em 36 (trinta e seis) meses, com um prazo de toler&acirc;ncia de 18 (dezoito) meses, consoante cl&aacute;usula quarta do contrato (fls.39). Logo o prazo final para entrega do im&oacute;vel se encerra em 02/08/2018, levando em considera&ccedil;&atilde;o o prazo de toler&acirc;ncia estabelecido no contrato.&Eacute; cedi&ccedil;o que a taxa de evolu&ccedil;&atilde;o de obra &eacute; devida desde a aprova&ccedil;&atilde;o do financiamento at&eacute; o t&eacute;rmino da obra.&nbsp;Portanto, se a obra atrasar, &eacute; devido o pagamento da referida taxa ao banco que financiou o im&oacute;vel, no caso, o Banco do Brasil, at&eacute; a sua conclus&atilde;o. Sendo certo que ocorrendo a mora da construtora requerida em rela&ccedil;&atilde;o &agrave; entrega do im&oacute;vel, a parte autora n&atilde;o pode ser penalizada com o pagamento de tal encargo.&nbsp;No entanto, no caso em deslinde, ainda n&atilde;o houve mora da construtora, haja vista que ainda n&atilde;o fora encerrado o prazo estimado para entrega do im&oacute;vel. Nesse ponto impende destacar que &eacute; legal a cl&aacute;usula contratual que prev&ecirc; a prorroga&ccedil;&atilde;o do prazo razo&aacute;vel para entrega do im&oacute;vel, considerando o princ&iacute;pio pacta sunt servanda.&nbsp;Desta feita, verifica-se a aus&ecirc;ncia de probabilidade do direito da parte autora, haja vista que, consoante dito alhures, a taxa de evolu&ccedil;&atilde;o de obra &eacute; devida at&eacute; a conclus&atilde;o da obra, bem como que n&atilde;o houve mora por parte da construtora demandada, haja vista que n&atilde;o houve, ainda, atraso na entrega do im&oacute;vel, tendo em vista a cl&aacute;usula que prev&ecirc; prazo de toler&acirc;ncia para entrega do im&oacute;vel.Ante o exposto, por considerar ausente a probabilidade do direito (art. 300 do CPC/15), INDEFIRO o pedido de tutela de urg&ecirc;ncia requestado.Inclua-se o feito em pauta de audi&ecirc;ncia de concilia&ccedil;&atilde;o. Cite-se a parte r&eacute; eintime-a desta decis&atilde;o, bem como para que compare&ccedil;a &agrave; audi&ecirc;ncia na data designada pelo Cart&oacute;rio, o que deve ser feito com anteced&ecirc;ncia m&iacute;nima de 20 dias.Intime-se os autores por advogado constitu&iacute;do (art. 334, &sect;3&ordm;, CPC/15).Dever&aacute; a parte r&eacute; ser advertida da possibilidade do art. 334, &sect;5&ordm;, bem como do termo inicial do prazo de contesta&ccedil;&atilde;o (art. 335).Fiquem as partes advertidas, ainda, de que o n&atilde;o comparecimento injustificado &agrave; audi&ecirc;ncia de concilia&ccedil;&atilde;o &eacute; considerado ato atentat&oacute;rio &agrave; dignidade da justi&ccedil;a e ser&aacute; sancionado com multa de at&eacute; dois por cento da vantagem econ&ocirc;mica pretendida ou do valor da causa (art. 334, &sect;8&ordm;).Publique-se. Intime-seMacei&oacute;, 10 de maio de 2018.Henrique Gomes de Barros TeixeiraJuiz de Direito<br/><b>Vencimento: </b>01/06/2018
		</span> 
		
	</td>
</tr>


				
					









	
	
		
	


<tr class="fundoClaro" style="">
	<td width="120" style="vertical-align: top">
		03/05/2018
	</td>
	<td width="20" valign="top" aria-hidden="true">
		
	</td>
	<td style="vertical-align: top; padding-bottom: 5px">
		 
		
			
			
				Conclusos
			
		 
		 
		<br /> 
		<span style="font-style: italic;">
			
		</span> 
		
	</td>
</tr>


				
					









	
	
		
	


<tr class="fundoEscuro" style="">
	<td width="120" style="vertical-align: top">
		02/05/2018
	</td>
	<td width="20" valign="top" aria-hidden="true">
		
	</td>
	<td style="vertical-align: top; padding-bottom: 5px">
		 
		
			
			
				Conclusos
			
		 
		 
		<br /> 
		<span style="font-style: italic;">
			
		</span> 
		
	</td>
</tr>


				
					









	
	
		
	


<tr class="fundoClaro" style="">
	<td width="120" style="vertical-align: top">
		02/05/2018
	</td>
	<td width="20" valign="top" aria-hidden="true">
		
	</td>
	<td style="vertical-align: top; padding-bottom: 5px">
		 
		
			
			
				Distribuído por Sorteio
			
		 
		 
		<br /> 
		<span style="font-style: italic;">
			
		</span> 
		
	</td>
</tr>


				
			</tbody>

		
		
	
</table>
		

	

	                
	                
	                
	                
	                
	            
	        
	            
	                
	                
	                
	                
	                
	                    



















<div style="padding-top: 10px;">
	
	
		
		
			









<table width="100%" border="0" cellspacing="0" cellpadding="0">
	<tr valign="top">
		<td height="21" valign="top" nowrap background="/cpopg/imagens/spw/fundo_subtitulo.gif"> 
			
				
					<h2 class="subtitle">
						Petições diversas
						
					</h2>
				
				
			
		</td>
		<td background="/cpopg/imagens/spw/fundo_subtitulo2.gif" width="90%" aria-hidden="true">
			<img src="/cpopg/imagens/spw/final_subtitulo.gif" width="16" height="20" tabindex="-1">
		</td>
	</tr>
</table>

		
	
</div>

<!-- Tabela de petições diversas -->
<table style="margin-left:15px; margin-top:1px;" align="center" width="98%" border="0" cellspacing="0" cellpadding="1"  >
	
		
			<thead>
			    <tr class="label">
			      <th width="140"  style="text-align:left">Data</th>
			      <th width="*" >Tipo</th>
			    </tr>
			    <tr class="fundoEscuro" height="2" aria-hidden="true">
					<td></td>
					<td></td>
				</tr>
			</thead>
			<tbody>
				
					
						<tr class="fundoClaro"> 
							<td width="140"  style="text-align:left">
								21/09/2018
							</td>
							<td width="*">
								Petição <br/>
								
							</td>
						</tr>
					
				
					
						<tr class="fundoEscuro"> 
							<td width="140"  style="text-align:left">
								21/09/2018
							</td>
							<td width="*">
								Petição <br/>
								
							</td>
						</tr>
					
				
					
						<tr class="fundoClaro"> 
							<td width="140"  style="text-align:left">
								21/09/2018
							</td>
							<td width="*">
								Petição <br/>
								
							</td>
						</tr>
					
				
					
						<tr class="fundoEscuro"> 
							<td width="140"  style="text-align:left">
								24/09/2018
							</td>
							<td width="*">
								Petição <br/>
								
							</td>
						</tr>
					
				
					
						<tr class="fundoClaro"> 
							<td width="140"  style="text-align:left">
								10/10/2018
							</td>
							<td width="*">
								Contestação <br/>
								
							</td>
						</tr>
					
				
					
						<tr class="fundoEscuro"> 
							<td width="140"  style="text-align:left">
								16/10/2018
							</td>
							<td width="*">
								Contestação <br/>
								
							</td>
						</tr>
					
				
					
						<tr class="fundoClaro"> 
							<td width="140"  style="text-align:left">
								26/11/2018
							</td>
							<td width="*">
								Impugnação à Contestação <br/>
								
							</td>
						</tr>
					
				
					
						<tr class="fundoEscuro"> 
							<td width="140"  style="text-align:left">
								26/11/2018
							</td>
							<td width="*">
								Impugnação à Contestação <br/>
								
							</td>
						</tr>
					
				
					
						<tr class="fundoClaro"> 
							<td width="140"  style="text-align:left">
								29/11/2018
							</td>
							<td width="*">
								Petição <br/>
								
							</td>
						</tr>
					
				
					
						<tr class="fundoEscuro"> 
							<td width="140"  style="text-align:left">
								05/12/2018
							</td>
							<td width="*">
								Petição <br/>
								
							</td>
						</tr>
					
				
					
						<tr class="fundoClaro"> 
							<td width="140"  style="text-align:left">
								11/02/2019
							</td>
							<td width="*">
								Petição <br/>
								
							</td>
						</tr>
					
				
					
						<tr class="fundoEscuro"> 
							<td width="140"  style="text-align:left">
								12/02/2019
							</td>
							<td width="*">
								Petição <br/>
								
							</td>
						</tr>
					
				
					
						<tr class="fundoClaro"> 
							<td width="140"  style="text-align:left">
								11/07/2019
							</td>
							<td width="*">
								Petição <br/>
								
							</td>
						</tr>
					
				
			</tbody>
		
		
	
</table>
<!--  Tabela de petições diversas -->


	                
	                
	                
	                
	                
	                
	            
	        
	            
	                
	                
	                
	                
	                    












<script type="text/javascript">
	(function($) {
		$(function(){
			var captcha = $.saj.getUrlParameter('uuidCaptcha');
			if(!captcha){
				return;
			}
			$('.incidente a').each(function(){
				var $incidente = $(this);
				var url = $incidente.attr('href');
				$incidente.attr('href', url+'&uuidCaptcha='+captcha);
			});
		})
	})(jQuery);
</script>










<div style="padding-top: 10px;">
	
	
		
		
			









<table width="100%" border="0" cellspacing="0" cellpadding="0">
	<tr valign="top">
		<td height="21" valign="top" nowrap background="/cpopg/imagens/spw/fundo_subtitulo.gif"> 
			
				
					<h2 class="subtitle">
						Incidentes, ações incidentais, recursos e execuções de sentenças
						
					</h2>
				
				
			
		</td>
		<td background="/cpopg/imagens/spw/fundo_subtitulo2.gif" width="90%" aria-hidden="true">
			<img src="/cpopg/imagens/spw/final_subtitulo.gif" width="16" height="20" tabindex="-1">
		</td>
	</tr>
</table>

		
	
</div>

<table style="margin-left:15px; margin-top:1px;" align="center" width="98%" border="0" cellspacing="0" cellpadding="1"  >

	
	
		<tr>
		  <td colspan="3" align="left">
		 	
 				
 				
 					Não há incidentes, ações incidentais, recursos ou execuções de sentenças vinculados a este processo.
 				
 			
		  </td> 
	    </tr>
	

</table>
<!--  Incidentes -->

	                
	                
	                
	                
	                
	                
	                
	            
	        
	            
	                
	                
	                
	                
	                
	                
	                
	                
	                
	                
	                    














<script type="text/javascript">
    (function ($) {
        $(function () {
            var captcha = $.saj.getUrlParameter('uuidCaptcha');

            if(!captcha){
                return;
            }
            $.each($('.processoApensado'), function (i, value) {
                var $link = $(value);
                $link.attr('href', $link.attr('href') + '&uuidCaptcha=' + captcha);
            })

        })
    })(jQuery);
</script>








<div style="padding-top: 10px;">
	
	
		
		
			









<table width="100%" border="0" cellspacing="0" cellpadding="0">
	<tr valign="top">
		<td height="21" valign="top" nowrap background="/cpopg/imagens/spw/fundo_subtitulo.gif"> 
			
				
					<h2 class="subtitle">
						Apensos, Entranhados e Unificados
						
					</h2>
				
				
			
		</td>
		<td background="/cpopg/imagens/spw/fundo_subtitulo2.gif" width="90%" aria-hidden="true">
			<img src="/cpopg/imagens/spw/final_subtitulo.gif" width="16" height="20" tabindex="-1">
		</td>
	</tr>
</table>

		
	
</div>



    
        <table style="margin-left:15px; margin-top:1px;" align="center" width="98%" border="0" cellspacing="0" cellpadding="1">
            <tbody id="dadosApensosNaoDisponiveis">
            <tr>
                <td colspan="3" align="left">Não há processos apensados, entranhados e unificados a este processo.</td>
            </tr>
            </tbody>
        </table>

    
    


	                
	            
	        
	            
	                
	                
	                
	                
	                
	                
	                
	                    






<div style="padding-top: 10px;">
	
	
		
		
			









<table width="100%" border="0" cellspacing="0" cellpadding="0">
	<tr valign="top">
		<td height="21" valign="top" nowrap background="/cpopg/imagens/spw/fundo_subtitulo.gif"> 
			
				
					<h2 class="subtitle">
						Audiências
						
					</h2>
				
				
			
		</td>
		<td background="/cpopg/imagens/spw/fundo_subtitulo2.gif" width="90%" aria-hidden="true">
			<img src="/cpopg/imagens/spw/final_subtitulo.gif" width="16" height="20" tabindex="-1">
		</td>
	</tr>
</table>

		
	
</div>

	                    












	
		
		    
				
			
			
						
	
	

<a name="audienciasPlaceHolder"></a>
<table style="margin-left:15px; margin-top:1px;" align="center" width="98%" border="0" cellspacing="0" cellpadding="1">
	
		
			<tr class="label"> 
				<th align="left" valign="top" width="140">Data</th>
				<th align="left" valign="top" width="*">Audiência</th>
				<th align="left" valign="top" width="100">Situação</th>
				<th align="left" valign="top" width="100">Qt. Pessoas</th>
			</tr>
			<tr class="fundoEscuro" height="2" aria-hidden="true">
				<td></td>
				<td></td>
				<td></td>
				<td></td>
			</tr>				    
			
				<tr class="fundoClaro"> 
					<td valign="top" align="left" valign="top" width="140">
						24/09/2018
					</td>
					<td valign="top" align="left" valign="top" width="*">
						
							
							
								Conciliação
							
											     
					</td>
					<td  align="left" valign="top" width="100">
						Realizada
					</td>
					<td  align="left" valign="top" width="100">
						7
					</td>
				</tr>        
				  
		
		
	
</table>
	                
	                
	                
	                
	            
	        
	        











	<table id="" class="secaoBotoesBody" width="100%" style="" cellpadding="2" cellspacing="0" border="0">
		<tr>
 			
			<td align="center">
			
	            
	            
	        
			</td>
		</tr>
	</table>


	    
	        










	
		











	
	
		



</td>
</tr>
</table>
<table width="100%" border="0" cellpadding="0" cellspacing="0" class="esajTabelaRodape" summary=" ">
    <tr>
        <td class="esajCelulaRodapeCentro">
            Desenvolvido pela Softplan em parceria com o Tribunal de Justiça de Alagoas
            
        </td>
    </tr>
</table>
<link href="https://www2.tjal.jus.br/esaj/tema/padrao/css/saj-popup-modal.css" rel="stylesheet" type="text/css">
<link href="https://www2.tjal.jus.br/esaj/tema/padrao/css/popup-beta.css" rel="stylesheet" type="text/css">


<script>
    var appEsajLayout = appEsajLayout || {};
    appEsajLayout.tituloCertificadoDigital = 'Certificado digital';
    appEsajLayout.cliente = 'AL';
    appEsajLayout.emailParaContato = '';
    appEsajLayout.utilizarBotaoContato = 'false';
    appEsajLayout.contexto = 'https://www2.tjal.jus.br/esaj';

    if(window.jQuery) {
        jQuery.saj = jQuery.saj || {};
        jQuery.saj.certificado = jQuery.saj.certificado || {};
        jQuery.saj.certificado.cliente = 'AL';
        jQuery.saj.certificado.tituloCertificadoDigital = 'Certificado digital';
        jQuery.saj.certificado.tituloProblemasAoAssinar = 'Falha de comunicação com o dispositivo de assinatura digital';
        jQuery.saj.certificado.urlConteudoAjudaWebSigner = appEsajLayout.contexto + '/ajudaWebSigner.do';
        jQuery.saj.certificado.licenca = 'AqYAKi5zYWphZG0udGpzcC5qdXMuYnIsKi5zZWEuc2MuZ292LmJyLCouc29mdHBsYW4uY29tLmJyLCoudGphYy5qdXMuYnIsKi50amFsLmp1cy5iciwqLnRqYW0uanVzLmJyLCoudGpjZS5qdXMuYnIsKi50am1zLmp1cy5iciwqLnRqcHIuanVzLmJyLCoudGpzYy5qdXMuYnIsKi50anNwLmp1cy5ickMAaXA0OjEwLjAuMC4wLzgsaXA0OjEyNy4wLjAuMC84LGlwNDoxNzIuMTYuMC4wLzEyLGlwNDoxOTIuMTY4LjAuMC8xNggAU3RhbmRhcmQAAAABcfpZUu/QD3s2bqedB8v9T7MiMQxcHNu0pUzGzO+7Ta++o+BbAVgPau7lzj6IUFI7RPFkPSnm0ZbCFAVBnKrqinyvE8MwEvmQhFcjlhGK+FfVV5mhVkkxSC3khDEUR8gsAPVvFncyyW2vTU2ODkyGdU0S/g62wHFtSwfGP6OFqlWpZTqbxbEGd1vfe9yHHulTYI7RcsGxaTMD6XgtPlStbOn9DxSYhN3S+4oTtIq7pUR6ETavPqYlXh5hwkxIpTe3Sf4awQucSigQ95sS9g7QpcGyVUVscm0J/rah9YwPCnHShzsttWA9aCmKrNdgFSTTo8TUOmESEP93TH6Z3dss+A==';
        jQuery.saj.certificado.cofreDigital = jQuery.saj.certificado.cofreDigital || {};
        jQuery.saj.certificado.cofreDigital.habilitado = 'true' === 'true';
        jQuery.saj.certificado.cofreDigital.enderecoBase = 'https://www2.tjal.jus.br/esaj';
    }

</script>
<script src="https://www2.tjal.jus.br/esaj/js/app-bundle.js?n=447896787"></script>
<script src="https://www2.tjal.jus.br/sajcas/seletorVersaoBeta.js?n=5f11e2fd-7ee6-4b83-8e39-bd75f17342b2&versao=1"></script>
<script language="javascript" type="text/JavaScript" src="https://www2.tjal.jus.br/esaj/js/softplan-websigner-2.5.1.js?n=736057767"></script>

<script language="javascript" type="text/JavaScript" src="https://www2.tjal.jus.br/esaj/js/saj-certificado.js?n=2205257179"></script>



<div id="webSignerNaoInstalado" style="display: none;">
    <p>Para utilização do certificado digital no Portal e-SAJ é necessário a instalação do plug-in Web Signer. Por favor, realize a instalação antes de continuar.</p>
    <div class="footer_certificado">
        <button id="redirecionarParaPaginaInstalacao" class="actionPrincipal spwBotaoDefault">Instalar</button>
        <button class="fecharModalInstalacao spwBotao">Cancelar</button>
        <button class="spwBotao" onclick="window.open('https://www2.tjal.jus.br/WebHelp//#id_instalacao_lacuna.htm','_blank')">Ajuda</button>
    </div>
</div>

<div id="webSignerNaoSuportado" style="display: none;">
    <p>O navegador utilizado não é compatível com a tecnologia (Web Signer) necessária para utilização do certificado digital. Por favor, utilize um dos navegadores homologados. Em caso de dúvidas, efetue a validação de requisitos <a href="https://www2.tjal.jus.br/petpg/abrirVerificacaoRequisitosPet.do" target="_blank">aqui</a>.</p>
    <div class="footer_certificado">
        <button class="fecharModalInstalacao spwBotaoDefault actionPrincipal">OK</button>
    </div>
</div>
	


  
	
	



	


	
	
	












	<form action="/cpopg/show.do?processo.codigo=01000O7550000&processo.foro=1&uuidCaptcha=sajcaptcha_e7177bac99504aa4b68783240bfba9d8" id="popupSenha" style="display: none;" method="POST">
		
		<table>
			
				
					<tr>
						<td colspan="2" style="padding-left: 20px" class="orientacao121" tabindex="2" aria-label="Atendendo o que está exposto na Res. 121 do CNJ.">
							Atendendo o que está exposto na Res. 121 do CNJ.
						</td>
					<tr>
					<tr>
						<td colspan="2" tabindex="3" aria-label="Será necessário informar uma senha para acessar processos em segredo de justiça, bem como para acessar autos dos demais processos. Caso não a possua e seja parte do processo, dirija-se ao cartório para solicitá-la.">
						<img src="/cpopg/imagens/setaTopico.gif">
						Será necessário informar uma senha para acessar processos em segredo de justiça, bem como para acessar autos dos demais processos. Caso não a possua e seja parte do processo, dirija-se ao cartório para solicitá-la.</td>
					</tr>
					<tr>
						<td colspan="2" tabindex="4" aria-label="Se for advogado (a) neste processo habilite-se no Portal ou efetue login pelo link "Identificar-se". O número de sua OAB no cadastro do Portal deverá ser igual ao número nos dados do processo.">
						<img src="/cpopg/imagens/setaTopico.gif">
						Se for advogado (a) neste processo habilite-se no Portal ou efetue login pelo link "Identificar-se". O número de sua OAB no cadastro do Portal deverá ser igual ao número nos dados do processo.</td>
					</tr>
				
				
			
			<tr align="center">
				<td align="right" width="47%"><b style="margin: 10px 0;">Senha do Processo:</b></td>
				<td align="left">
					<input type="password" name="senhaProcesso" maxlength="12" value="" rotulo="Senha do Processo" style="margin: 10px 0;" id="senhaProcesso">
					<input type="hidden" name="idRecursoQueProvocouLiberacaoPorSenha" value="">
				</td>
			</tr>
			<tr>
				<td colspan="2" align="center">
					
						
						
							













	
		
	
	<input type="submit" name="btEnviarSenha" value="Continuar" onclick="" onmouseover="B_mOver(this);" onmouseout="B_mOut(this);" onmouseover="B_mOver(this);" onmouseout="B_mOut(this);" class="spwBotaoDefault " id="btEnviarSenha">

&nbsp;
						
					
					

















	
	
	    
	<input type="button" name="btFechar" value="Fechar" onclick="" onmouseover="B_mOver(this);" onmouseout="B_mOut(this);" class="spwBotao " id="botaoFecharPopupSenha">


				</td>
			</tr>
		</table>
	</form>



	
	
</div>

<script type="text/javascript">window.NREUM||(NREUM={});NREUM.info={"errorBeacon":"bam.nr-data.net","licenseKey":"b61bdf610d","agent":"","beacon":"bam.nr-data.net","applicationTime":76,"applicationID":"111652883","transactionName":"MlZRN0QECkMAVERZDgscYBdEEBBDIFREWQ4LHFERGAYLXU9EX1YVFV9SDRgWBVpPVEBfTxZHQRZCFkpRAkNZXw9LYFsMQSQHRAhYXg==","queueTime":0}</script></body>
</html>
`