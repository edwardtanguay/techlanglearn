# L'essentiel de Elasticsearch

https://www.linkedin.com/learning/l-essentiel-de-elasticsearch

- duration: 01:59:00
- language: fr
- topics: elasticsearch
- rank: 4.8
- description: good overview, in French
- year: 2022

## Bienvenue dans « L'essentiel d'Elasticsearch » , 0:37, 2025-02-24

- intro

## Découvrir Elasticsearch, 2:44, 2025-02-24

- project "Elastic" known before as "ELK"
- Elasticsearch, Logstash, and Kibana
- based on REST
- "Elastic Stack" is the new term for "ELK stack" (2025-02-24) and offers: Kibana, Elasticsearch, Integrations
- the company Elastic has expanded the stack’s capabilities beyond just those three core tools
- they've added other products like Beats (lightweight data shippers) and a wide range of cloud solutions
- but Logstash is still a central part of most of the implementations

## Comprendre l'utilisation d'Elasticsearch, 4:05, 2025-02-24

- advantage of Elasticsearch over database
  - when you need fast, flexible, and powerful search capabilities
  - fast for full-text searches, uses inverted indexes
  - fuzzy search, partial matching, phrase matching, relevance scoring, and proximity searches
  - built to handle large volumes of data
  - Elasticsearch indexes documents in near real-time

## Installer et démarrer son premier nœud Elasticsearch, 4:20, 2025-02-24

- he installed on Debian directly
- it runs on 9300
  - installs a cluster here
- then he talks about port 9200
  - this is the REST/JSON port
  - does `curl localhost:9200`
- he then changes .yml settings regarding cluster

## Démarrer un cluster avec trois nœuds, 2:01, 2025-02-24

- installing a cluster and three nodes
  - starts the third node in the same way
  - then has a cluster with three nodes

## Démarrer son instance Kibana, 3:04, 2025-02-24

- installing a cluster and three nodes
  - starts the third node in the same way
  - then has a cluster with three nodes##withtrehenods
- he is now installing Kibana

## Définir un document dans Elasticsearch, 4:43, 2025-02-24

https://www.linkedin.com/learning/l-essentiel-de-elasticsearch/definir-un-document-dans-elasticsearch?autoSkip=true&resume=false

- what is a document
  - an object in JSON
  - can have nested arrays and objects

## Indexer ses premiers documents, 4:10, 2025-02-24

- when you put an object, it breaks up the words of strings and indexes them

## Faire ses premiers pas avec le mapping, 4:14, 2025-02-24

https://www.linkedin.com/learning/l-essentiel-de-elasticsearch/faire-ses-premiers-pas-avec-le-mapping?autoSkip=true&resume=false

- mapping
  - elasticsearch can't modify an existing object
  - the parameters have to be the same type
  - for example, if you enter a date for a field, that field will be considered date type
  - you can recreate the type mapping, i.e. a new index

## Utiliser les API CRUD, 5:05, 2025-02-25

https://www.linkedin.com/learning/l-essentiel-de-elasticsearch/utiliser-les-api-crud?autoSkip=true&resume=false

- CRUD
  - (click arrow or CTRL-Enter)
  - GET /person/\_doc/2
    - in modern versions of Elasticsearch, document types are deprecated, so this is more of a placeholder
    - but you still have to include it, even in 8.17.2
  - DELETE /person/\_doc/2
  - PUT /person/\_doc/2 (with object) creates or replaces
    - I get "unknown field" with this
    - in the course it worked
  - POST /person/\_doc
    - it creates an id
    - each time it creates a new id##crudexamplkes

## Découvrir l'API Bulk, 3:47, 2025-02-25

- DELETE logs, metrics
- indexing
- you can index many at a time##indexmenay

## Chercher des documents, 4:46, 2025-02-25

- multi_match and fuzziness

## Comprendre l'analyse de texte, 2:49, 2025-02-25

- features for better search-UX in various languages

## Créer son propre analyseur, 4:42, 2025-02-25

- tokenizer: "whitespace"

## Créer un mapping, 3:10, 2025-02-25

- you can output a mapping, then edit it to replace it with some changes

## Chercher dans ses différents champs, 3:08, 2025-02-25

- it can do ranges as well

## Utiliser le champ de type Runtime, 4:29, 2025-02-25

- runtime enables a kind of computed values to search

## Regrouper les documents par champ, 3:44, 2025-02-25

https://www.linkedin.com/learning/l-essentiel-de-elasticsearch/regrouper-les-documents-par-champ?autoSkip=true&resume=false

- regrouping
  - easily count based on field

## Utiliser les agrégations multi-niveaux, 5:00, 2025-02-25

- multi-level agregations

## Combiner recherche et agrégation, 2:59, 2025-02-25

- runtime is slower than indexing

## Utiliser la navigation par facettes, 4:40, 2025-02-25

- he combines features and does more intricate queries

## Définir des champs géographiques dans le mapping, 2:28, 2025-02-25

- geopoints
  - various formats

## Rechercher des documents dans un périmètre donné, 3:44, 2025-02-25

- he saves a polygon and then searches if someone is in the polygon

## Trouver les documents les plus proches, 2:42, 2025-02-25

- can find locations that are within a certain range
- filter before for better performance

## Utiliser les agrégations sur des coordonnées géographiques, 5:03, 2025-02-25

- finds if people live within 100m, 200m, 300m of a location##thelocccksi
- also has geo_centroid and geo_bounds

## Modifier les documents à la volée avec l'API Ingest, 5:24, 2025-02-25

- change data on the fly
  - with the API \_ingest
  - with a pipeline that changes values##wihtpipeplin
  - some documents are in the pipeline, some are not

## Comprendre quelques processeurs Ingest essentiels, 3:26, 2025-02-25

- grok is complex
- example of dissect

## Découvrir les processeurs Fingerprint et Geoip, 4:34, 2025-02-25

- fingerprint
  - creates a unique hash fingerprint for a document or specific fields, is used to deduplicate data or create consistent identifiers for similar content
- geoip
  - enriches documents with geolocation data based on an IP address
  - uses MaxMind’s GeoLite2 database behind the scenes

## Indexer son document PDF, 3:08, 2025-02-25

- Tika
- GET /\_cat/plugins?v

## Réindexer ses documents, 2:51, 2025-02-25

- reindexing
  - talks about slices

## Réindexer ses documents via un pipeline, 4:16, 2025-02-25

- you can reindex on the fly

## Conclure ce cours sur Elasticsearch, 3:38, 2025-02-25

- you can create an index in Kibana
- you can view the data in Kibana##withtdashs

## VOCAB - FRENCH

```
this may prove useful
cela peut s'avérer utile
2025-02-26 01:01:23

basically
en gros
2025-02-26 00:49:33

apart from the date
mise à part la date
2025-02-26 00:47:40

the convert processor will help us
le processeur convert va nous aider
2025-02-26 00:40:21

we wish to make it
nous souhaitons en faire
2025-02-26 00:38:40

opening and closing bracket
crochet ouvrant et fermant
2025-02-26 00:35:12

Let's focus on some processors
attardons-nous sur quelques processeurs
2025-02-26 00:31:48

whatever it was
quoi que ce soit
2025-02-26 00:28:10

alleviate
alléger
2025-02-26 00:19:04

such as a return code
telles qu'un code retour
2025-02-26 00:16:17

upper left
superior gauche
2025-02-26 00:04:23

we are on a scale of a few centimeters
on est à l'échelle de quelques centimètres
2025-02-25 23:58:43

as I go, I will withdraw the request
en passant, je vais retirer la requête
2025-02-25 23:54:35

all or part of our
tout ou partie de nos
2025-02-25 23:53:00

are distributed
sont réparties
2025-02-25 23:51:29

rings
des anneaux
2025-02-25 23:50:37

widely used on platforms
très usités sur les plateformes
2025-02-25 17:37:30

to refine the results
pour affiner les résultats
2025-02-25 17:36:18

Elasticsearch querying allows fast document searches
le requêtage Elasticsearch permet de rechercher des documents rapidement
2025-02-25 17:33:44

on the fly
à la volée
2025-02-25 17:31:59

and so on
ainsi de suite
2025-02-25 17:28:25

they are distributed by day
elles sont répartie par jour
2025-02-25 17:24:04

as we saw earlier
comme nous l'avions vu toute à l'heure
2025-02-25 17:16:19

this is done with
cela se fait avec
2025-02-25 17:07:25

distribution of our documents
répartition de nos documents
2025-02-25 17:05:56

however
pour autant
2025-02-25 12:09:59

to overcome this problem
pour pallier ce problème
2025-02-25 11:05:48

side effects
des effets de bords
2025-02-25 11:02:19

completely deleted
carrément supprimé
2025-02-25 10:58:58

from a business point of view
d'un point de vue métier
2025-02-25 10:45:05

to get down to business
de s'atteler à la tâche
2025-02-25 10:05:54

we realize that
on se rende compte que
2025-02-25 09:10:17

indexing one by one can be very expensive
indexer un par un peu peut s'avérer très coûteux
2025-02-25 09:02:52

some bandwidth
un peu de bande passante
2025-02-25 08:35:18

which we can also verify
ce qu'on peut d'ailleurs vérifier
2025-02-25 08:20:59

as well as data
ainsi que des données
2025-02-25 08:19:21

to retrieve a document from the datastore
pour récupérer un document de la base de données
2025-02-25 08:05:02

to sort
pour faire du tri
2025-02-25 00:54:59

in fact
en fait;pr=en fett
2025-02-25 00:50:37

a nested object
un objet imbriqué
2025-02-24 23:03:51

once the installation is complete
une fois l'installation achevée
2025-02-24 22:19:25

the same in a search engine
idem dans un moteur de recherche
2025-02-24 20:58:12

since it is sorted alphabetically
puisqu'il est trié de façon alphabétique
2025-02-24 20:56:16

did he tidy up my data
a-t-il rangé ma donnée
2025-02-24 20:50:15

the expected results
les résultats escomptés
2025-02-24 20:48:35

and typos
et les fautes de frappe
2025-02-24 20:46:19

we are lucky because
on a de la chance puisque
2025-02-24 20:44:43

what will have to be done
ce qu'il va falloir faire
2025-02-24 20:41:14

if our user types "cd"
si notre utilisateur saisit "cd"
2025-02-24 20:37:32

a way to get out
une façon de s'en sortir
2025-02-24 20:35:27

it doesn't bring any hits
ça ne ramène aucun hit; conj=ramener; ramener = to bring back
2025-02-24 20:32:57

compared to a search
par rapport à une recherche
2025-02-24 20:29:30

an integration layer
une couche d'intégration
2025-02-24 10:52:00

as part of this discovery
dans le cadre de cette découverte
2025-02-24 10:26:35

to block them before they spread
de les bloques avant qu'elles ne se propagent
2025-02-24 10:25:24

manage more data
gérer davantage de données
2025-02-24 10:23:10

materials made available to him
matérielles mises à sa disposition; conj=mettre; mises = past participle (feminine plural)
2025-02-24 09:32:57

extend vertically
s'étendre verticalement
2025-02-24 09:30:48

on a data field
sur un champ données
2025-02-24 09:28:58

research on value ranges
de la recherche sur des plages de valeur
2025-02-24 09:26:05

and regardless of the source
y ce quelle qu'en soit la source
2025-02-24 09:10:23

which will allow us to search
qui nous permettront de rechercher; conj=permettre, permettront = third-person future
2025-02-24 09:06:30

timestamped
horodatées
2025-02-24 09:04:35

allows you to dig into their data
permet de fouiller dans leurs données
2025-02-24 09:03:52

```
