This library is a fork of "abhiyerra/hubspot"
with other new functionalities for contacts data retrive and other stuffs.

#+TITLE: HubSpot

HubSpot API in Go.


#+begin_src go
apiKey := os.Getenv("HUBSPOT_API_KEY")

a := NewContact(apiKey, "abhi@acksin.com")
a.Add("firstname", "Abhi")
a.Add("lastname", "Yerra")
a.Add("company", "Acksin")
a.Add("lifecyclestage", "opportunity")
a.Add("acksinsoftware", "opsZero")
resp := a.Publish()
if resp.Vid != 901 {
        t.Errorf("Failed to update contact")
}

d := NewDeal(apiKey)
d.Associations.AssociatedVids = []int{resp.Vid}
d.Add("dealname", "Tim's Newer Deal")
d.Add("dealstage", "closedwon")
d.Add("closedate", Timestamp())
d.Add("amount", "60000")
d.Add("dealtype", "newbusiness")
d.Publish()
#+end_src
