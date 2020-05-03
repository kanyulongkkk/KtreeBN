
/* enum for scores, to be matched from the label string passed down from R. */
typedef enum {
  ENOSCORE  =   0, /* error code, no such score. */

  LOGLIK    =   1, /* log-likelihood, discrete data. */
  AIC       =   2, /* AIC, discrete data. */
  BIC       =   3, /* BIC, discrete data. */
  BDE       =   4, /* Bayesian Dirichlet equivalent score. */
  BDS       =   5, /* Bayesian Dirichlet sparse score. */
  BDJ       =   6, /* Bayesian Dirichlet with Jeffrey's prior. */
  K2        =   7, /* K2 score. */
  MBDE      =   8, /* Bayesian Dirichlet equivalent score, interventional data .*/
  BDLA      =   9, /* Bayesian Dirichlet score, locally averaged. */

  LOGLIK_G  = 100, /* log-likelihood, Gaussian data. */
  AIC_G     = 101, /* AIC, Gaussian data. */
  BIC_G     = 102, /* BIC, Gaussian data. */
  BGE       = 103, /* Bayesian Gaussian equivalent score. */

  LOGLIK_CG = 200, /* log-likelihood, conditional Gaussian data. */
  AIC_CG    = 201, /* AIC, conditional Gaussian data. */
  BIC_CG    = 202, /* BIC, conditional Gaussian data. */
} score_e;

score_e score_label(const char *label);

/* enum for graph priors, to be matched from the label string passed down from R. */
typedef enum {
  ENOPRIOR  =  0, /* error code, no such graph prior. */
  UNIFORM   =  1, /* uniform prior. */
  VSP       =  2, /* variable selection prior. */
  CS        =  3, /* Castelo & Siebes prior. */
  MU        =  4, /* marginal uniform prior. */
} gprior_e;

gprior_e gprior_label(const char *label);

/* score delta from score.delta.c */
SEXP score_delta(SEXP arc, SEXP network, SEXP data, SEXP score,
    SEXP score_delta, SEXP reference_score, SEXP op, SEXP extra, SEXP decomposable);

/* from graph.priors.c */
double graph_prior_prob(SEXP prior, SEXP target, SEXP beta, SEXP cache,
    int debuglevel);

/* from per.node.score.c */
SEXP per_node_score(SEXP network, SEXP data, SEXP score, SEXP targets,
    SEXP extra_args, SEXP debug);
void c_per_node_score(SEXP network, SEXP data, SEXP score, SEXP targets,
    SEXP extra_args, int debuglevel, double *res);

/* score functions exported to per.node.score.c */
double dlik(SEXP x, double *nparams);
double cdlik(SEXP x, SEXP y, double *nparams);
double loglik_dnode(SEXP target, SEXP x, SEXP data, double *nparams, int debuglevel);
double glik(SEXP x, double *nparams);
double cglik(SEXP x, SEXP data, SEXP parents, double *nparams);
double c_fast_ccgloglik(double *xx, double **gp, int ngp, int nobs, int *config,
    int nconfig);
double loglik_gnode(SEXP target, SEXP x, SEXP data, double *nparams, int debuglevel);
double loglik_cgnode(SEXP target, SEXP x, SEXP data, double *nparams, int debuglevel);
double dirichlet_node(SEXP target, SEXP x, SEXP data, SEXP iss, int per_node,
    SEXP prior, SEXP beta, SEXP experimental, int sparse, int debuglevel);
double dirichlet_averaged_node(SEXP target, SEXP x, SEXP data, SEXP l,
    SEXP prior, SEXP beta, int sparse, int debuglevel);
double wishart_node(SEXP target, SEXP x, SEXP data, SEXP isize, SEXP phi,
    SEXP prior, SEXP beta, int debuglevel);

