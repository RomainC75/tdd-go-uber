# UberLike SAAS!

L'application Uber-Like qui permet de réserver un Driver instantanément !

## **User Story - Réserver une course**

En tant que **Rider**,  
Je souhaite **réserver une course** pouvant m'amener à ma destination  
De sorte à assurer une alternative efficace aux transports en commun.

Un Rider peut réserver une course à tout moment pour n'importe quelle destination.  
La réservation est confirmée une fois qu'un Driver libre est assigné par le système.

---

## **Règles de Gestion de la Tarification**

### **1. Prix de Base par Direction**
- **Paris vers l'extérieur** : 20€.
- **Paris intra-muros (trajet à l'intérieur de Paris)** : 30€.
- **Extérieur vers Paris** : 10€.
- **Extérieur vers l'extérieur** : 50€.

### **2. Frais Kilométriques**
- Si forfait Basic, un supplément de **0,5€ par kilomètre** parcouru est ajouté au prix de base du trajet.
- Si forfait Premium, un supplément de **0.5€ par kilomètre** parcouru est ajouté au prix de base du trajet **à partir de 5km parcourus**.
- Les frais kilométriques sont calculés uniquement sur la distance parcourue,  
  sans tenir compte des bouchons ou autres ralentissements.

### **3. Changement de Trajet**
- Tout **changement de trajet** en cours de route entraîne un **surcoût de 10€**,  
  quelle que soit la formule initiale. 

### **4. Option UberX**
- **Description** : UberX permet d'accéder à des véhicules de type berline haut de gamme.
- **Supplément** : Un supplément de **10€** est ajouté pour bénéficier du mode UberX.
- **Condition de Distance** : Le mode UberX est disponible uniquement pour les trajets de **3 kilomètres ou plus**.  
  Pour les trajets inférieurs à cette distance, l'option est refusée.
- **Offre Anniversaire** : Si le jour de la course correspond à l'anniversaire du Rider,  
  le supplément de 10€ pour l'option UberX est **offert** pour les trajets de 3 kilomètres ou plus.

### **5. Offre de Bienvenue**
- Une remise de **X €** est offerte aux Riders durant leur première année d'utilisation du service.

### **6. Restriction sur les Réservations**
- Un Rider **ne peut réserver qu'une seule course à la fois**.
- Il est impossible de réserver une nouvelle course tant que la course en cours n'est pas terminée ou annulée.

---

## **Règles d'assignation d'un Driver**
- Le système recherche un Driver disponible dans un rayon de **5 km** autour du Rider.
- Un Driver ne peut être assigné qu'à **une seule réservation** à la fois.

---

## **User Story - Lister toutes mes courses passées**
En tant que **Rider**,  
Je souhaite **lister tout l'historique de mes courses avec mention des Drivers respectifs**  
De sorte à pouvoir me figurer la fréquence de mon utilisation.
