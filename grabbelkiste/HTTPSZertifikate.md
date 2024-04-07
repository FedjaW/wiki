# HTTPS-Zertifikate für Server

An SSL certificate is part of the TLS protocol.
Wofür HTTPS benötingt wird:

1. Verschschlüsselung zwischen dem Datenverkehr zwischen Client und Server
   Verfahren: AES (Advanced Encryption Standard) -> Symmetrisches Verfahren
   Dafür braucht man einen gmeinsamen Schlüssel. Der Schlüssel wird ausgetauscht mit dem Diffy-Hellman Verfahren, man kennt den im Vorfeld nicht.
   Bis hierhin hat es noch nichts mit Zertifikaten zu tun.

2. Die Identität des Servers für den Client Überprüfbar zu machen
   Der Client muss checken dass er mit dem richtigen Server spricht.
   Dafür wird ein Asymmetrisches Verschlüsselungsverfahren genutzt. (Public/Private Key)
   Der Server kennt seinen privaten Schlüssel und der Client kennt den öffentlichen Schlüssel der Servers
   Da nur der Server den privaten Schlüssel kennt (so die Hoffnung ;-) dann kann der Client sicherstellen dass er mit dem richtigen Server spricht, ansonsten würde die verschlüsselte Kommunikation nicht funktionieren weil die Schlüssel nicht zusammenpassen.
   Der zu dem private Schlüssel des Servers zugehörige öffentliche Schlüssel irgendwie zum Client kommen muss, und zwar auf einem Weg mit dem der Client überprüfen kann dass es sich bei dem öffentlichen Schlüssel tatsächlich um den offiziellen Schlüssel des gewünschten Servers handelt.
   Also: Das System basiert darauf, feststellen zu können, den zum echten/richtigen Server passenden (öffentlichen) Schlüssel, in den Händen zu halten.

   Was man also macht: Man stattet den öffentlichen Schlüssel mit Metadaten aus.
   Die wichtigste Information ist der Domain Name zu dem der öffentlichen Schlüssel gehören soll. Das heißt man schafft eine künstliche Bindung von einem öffentlichen Schlüssel an eine Domain.
   So weiß der Client für welche Domain ein öffentlichen Schlüssel zuständig ist.

   Woher weiß der Client dass die Metadaten korrekt sind? (Theoretisch kann ja irgendjemand diese Metadaten hineinschreiben)
   Deshalb werden diese Metadaten zusammen mit einem "Fingerabdruck" des öffentlichen Schlüssels noch einmal von einer vertrauenswürdigen dritten Partei digital signiert.
   Die Idee ist also dass ich mich als Anwender auf diese dritte Partei verlassen kann.
   Das ganze steht und fällt mit dem Vertrauen in diese dritte Partei.

   Ein solcher öffentlichen Schlüssel der mit Metadaten versehen und zusätzlich auch signiert ist, nennt sich nun Certificate.
   Die vertrauenswürdige dritte Partei nennt sich nun Certificate-Authority (CA).

Grundlegender Ablauf um einen Server/Service mit einem HTTPS-Zertifikat auszustatten:

1. Man erstellt einen privaten und einen öffentlichen Schlüssel
2. Dann fügt man dem öffentlichen Schlüssel die Metadaten hinzu und
3. schickt das zusammen mit Unterlagen zur Verifikation der Identität an eine CA und von der erhält man dann ein signiertes Zertifikat

   Man hat es also mit vier Dokumenten zu tun:

   1. Den privaten Schlüssel
   2. Den öffentlichen Schlüssel
   3. Den Antrag für ein Zertifikat wo die Metadaten enthalten sind
   4. Das Zertifikat ansich

Detaillierte Ablauf um einen Server/Service mit einem HTTPS-Zertifikat auszustatten:

1. Um den privaten Schlüssel zu erzeugen wird unter der Haube Open-SSL verwendet. (SSL - Secure Sockets Layer)
   RSA Schlüssel Befehl: $ openssl genrsa -out privateKey.pem 2048
   2048 gibt die Schlüssellänge in Bit an, hier also 2048 Bit
2. Der öffentlichen Schlüssel muss aus der erzeugten .pem Datei extrahiert werden. Die Informationen für den öffentlichen Schlüssel liegen darin vor, jedoch in einer Datei.
   Um den öffentlichen Schlüssel zu bekommen: Befehl: $ openssl rsa -in privateKey.pem -pubout
3. Man braucht die oben genannten Befehle eher selten, da man alles in einem Abwasch machen kann mit diesem Befehl:
   $ openssl req -out csr.pem -new -newkey resa:2048 -nodes -keyout privateKey.pem
   csr steht fuer: certificate signing request
   die csr.pem Datei ist mein Antrag zum signieren, sie wird an die CA geschickt
4. Dann kommt das Zertifikat nach einer (unbestimmten (längeren)) Zeit zurück, idealerweise als certificate.pem (oder .certificate.key)
5. Danach braucht man die csr.pem (signing request datei) nicht mehr und kann sie löschen
6. Überbleibt die privateKey.pem und die certificate.pem -> Diese beiden kann (MUSS) man in seinen Service einbinden!
7. Das Einbinden der beiden Dateien ist von Technology zu Technology unterschiedlich.
   In node genügt es die beiden Dateien in das jeweilige Projekt zu kopieren, sie auszulesen und den jeweiligen Inhalt als Parameter an die HTTPS.createServer Funktion zu übergeben

Wichtig zu wissen:

1. Die von den CA ausgestellten Zertifikate haben ein Ablaufdatum
   Das Zertifikat muss von Zeit zu Zeit ausgetauscht werden
2. Das ganze funktioniert wie beschrieben, dafür braucht man aber eine registrierte Domain
3. Zertifikate von CA's kosten Geld
4. Für interne Services kann man selbst signierte Zertifikate nutzen
5. Lets Encrypt (eine CA) bieten eine API an um den Vorgang zu automatisieren
   Man braucht auf dem Server ein Skript das bevor das Zertifikat ausläuft einen neuen privaten Schlüssel und einen csr erzeugt, den an die CA schickt, das Zertifikat zurück erhält
   Die Verifikation geht schnell, nämlich ueber DNS Einträge
6. Lets Encrypt ist umsonst!
7. Fun Fact: Keine CA würde jemals ein Zertifikat für localhost ausstellen
   localhost kann man nur selbst signieren, aber niemals als offizielle Variante
