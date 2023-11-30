function setCookie(cookieName,cookieValue) {
    var today = new Date();
    const epoch = today.getTime()
    var expire = new Date(epoch + 60000); // 1 min
    document.cookie = cookieName+"="+escape(cookieValue)+ ";expires="+expire.toGMTString();
}

function getCookie(cname) {
    let name = cname + "=";
    let decodedCookie = decodeURIComponent(document.cookie);
    let ca = decodedCookie.split(';');
    for(let i = 0; i <ca.length; i++) {
      let c = ca[i];
      while (c.charAt(0) == ' ') {
        c = c.substring(1);
      }
      if (c.indexOf(name) == 0) {
        return c.substring(name.length, c.length);
      }
    }
    return "";
}

export {setCookie, getCookie};